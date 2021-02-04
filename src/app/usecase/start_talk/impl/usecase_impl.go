// Package impl ユースケース実装
package impl

import (
	"github.com/go-server-dev/src/app/usecase/start_talk"
	"time"
)

// UseCaseImpl ユースケース実装
type UseCaseImpl struct {
	readOnlyRepository start_talk.ReadOnlyRepository
	presenter          start_talk.Presenter
}

// インターフェースを満たしているかのチェック
var _ start_talk.UseCase = (*UseCaseImpl)(nil)

// NewUseCaseImpl ゲーム参加ユースケースインスタンスを新規作成する
func NewUseCaseImpl(
	readOnlyRepository start_talk.ReadOnlyRepository,
	presenter start_talk.Presenter) *UseCaseImpl {
	return &UseCaseImpl{
		readOnlyRepository: readOnlyRepository,
		presenter:          presenter,
	}
}

// Excute ゲームに参加する
func (j *UseCaseImpl) Excute(input start_talk.Input) {

	//出力DTO作成
	output := start_talk.Output{
		MemberID:      input.MemberID,
		GroupRoomID:   input.GroupRoomID,
		GroupRoomType: input.GroupRoomType,
		ReplyToken:    input.ReplyToken,
	}

	// ゲームマスターを取得
	gameMaster, err := j.readOnlyRepository.FindGameMasterByGroupID(input.GroupRoomID)
	if err != nil {
		output.Err = err
		j.presenter.Execute(output)
		return
	}

	// 既にトーク中の場合、「トーク中」エラー返却
	if gameMaster.IsTalkTime() {
		output.Err = start_talk.ErrAlreadyStarted
		j.presenter.Execute(output)
		return
	}

	// 状態をトークフェーズに設定
	gameMaster.StartTalk()

	// トーク時間を返却
	talkTime := gameMaster.TalkTimeMin()
	output.TalkTimeMin = talkTime

	// タイマーをセット
	// TODO：LINEグループ数によってインスタンスが増大するのが課題
	// ->対応案：送信日時をDBに保存してバッチ処理にするなど対策する必要がある
	// TODO：トーク時間を延長・終了した場合の処理を考慮する必要がある
	go setFinishTimer(j, input, talkTime)

	// 出力
	output.TalkTimeMin = talkTime
	output.Err = nil
	j.presenter.Execute(output)
}

// setFinishTimer 一定時間後、終了告知する
func setFinishTimer(j *UseCaseImpl, input start_talk.Input, duration time.Duration) {
	finishTalkOutput := start_talk.FinishTalkOutput{
		GroupRoomID:   input.GroupRoomID,
		GroupRoomType: input.GroupRoomType,
		ReplyToken:    input.ReplyToken,
	}
	timer := time.NewTimer(duration)
	<-timer.C
	j.presenter.FinishTalk(finishTalkOutput)
}
