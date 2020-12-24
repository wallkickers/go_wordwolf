// Package impl ユースケース実装
package impl

import (
	"github.com/go-server-dev/src/app/usecase/repository"
	"github.com/go-server-dev/src/app/usecase/start_talk"
)

// UseCaseImpl ユースケース実装
type UseCaseImpl struct {
	gameMasterRepository repository.GameMasterRepository
	readOnlyRepository   join.ReadOnlyRepository
	presenter            join.Presenter
}

// インターフェースを満たしているかのチェック
var _ start_talk.UseCase = (*UseCaseImpl)(nil)

// NewUseCaseImpl ゲーム参加ユースケースインスタンスを新規作成する
func NewUseCaseImpl(
	gameMasterRepository repository.GameMasterRepository,
	readOnlyRepository join.ReadOnlyRepository,
	presenter join.Presenter) *UseCaseImpl {
	return &UseCaseImpl{
		gameMasterRepository: gameMasterRepository,
		readOnlyRepository:   readOnlyRepository,
		presenter:            presenter,
	}
}

// Excute ゲームに参加する
func (j *UseCaseImpl) Excute(input start_talk.Input) join.Output {

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

	// 参加メンバーを追加
	talkTime = gameMaster.TalkTimeMin()
	if err = j.gameMasterRepository.Save(gameMaster); err != nil {
		output.Err = err
		j.presenter.Execute(output)
		return output
	}

	// 出力
	output.Err = nil
	j.presenter.Execute(output)
	return output
}
