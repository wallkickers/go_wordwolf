// Package impl ユースケース実装
package impl

import (
	"github.com/go-server-dev/src/app/usecase/accept_votes"
	"github.com/go-server-dev/src/app/usecase/repository"
)

// UseCaseImpl 投票するユースケース実装
type UseCaseImpl struct {
	gameMasterRepository repository.GameMasterRepository
	readOnlyRepository   accept_votes.ReadOnlyRepository
	presenter            accept_votes.Presenter
}

// インターフェースを満たしているかのチェック
var _ accept_votes.UseCase = (*UseCaseImpl)(nil)

// NewUseCaseImpl ゲーム参加ユースケースインスタンスを新規作成する
func NewUseCaseImpl(
	gameMasterRepository repository.GameMasterRepository,
	readOnlyRepository accept_votes.ReadOnlyRepository,
	presenter accept_votes.Presenter) *UseCaseImpl {
	return &UseCaseImpl{
		gameMasterRepository: gameMasterRepository,
		readOnlyRepository:   readOnlyRepository,
		presenter:            presenter,
	}
}

// Excute 投票する
func (j *UseCaseImpl) Excute(input accept_votes.Input) accept_votes.Output {

	//出力DTO作成
	output := accept_votes.Output{
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

	// 投票を管理
	gameMaster.SetVoteManagement(input.FromMemberID, input.ToMemberID)
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
