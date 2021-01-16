// Package impl ユースケース実装
package impl

import (
	"github.com/go-server-dev/src/app/usecase/start_talk"
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
func (j *UseCaseImpl) Excute(input start_talk.Input) start_talk.Output {

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
		return output
	}

	// 既にトーク中の場合、「トーク中」エラー返却
	if gameMaster.IsTalkTime() {
		output.Err = start_talk.ErrAlreadyStarted
		j.presenter.Execute(output)
		return output
	}

	// 状態をトークフェーズに設定
	gameMaster.StartTalk()

	// トーク時間を返却
	talkTime := gameMaster.TalkTimeMin()
	output.TalkTimeMin = talkTime

	// TODO：タイマーをセットする必要がある

	// 出力
	output.Err = nil
	j.presenter.Execute(output)
	return output
}
