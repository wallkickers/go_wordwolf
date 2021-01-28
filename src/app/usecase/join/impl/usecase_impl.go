// Package impl ユースケース実装
package impl

import (
	"github.com/go-server-dev/src/app/usecase/join"
	"github.com/go-server-dev/src/app/usecase/repository"
)

// UseCaseImpl ゲーム参加ユースケース実装
type UseCaseImpl struct {
	gameMasterRepository repository.GameMasterRepository
	readOnlyRepository   join.ReadOnlyRepository
	presenter            join.Presenter
}

// インターフェースを満たしているかのチェック
var _ join.UseCase = (*UseCaseImpl)(nil)

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
func (j *UseCaseImpl) Excute(input join.Input) join.Output {

	//出力DTO作成
	output := join.Output{
		MemberID:      input.MemberID,
		GroupRoomID:   input.GroupRoomID,
		GroupRoomType: input.GroupRoomType,
		ReplyToken:    input.ReplyToken,
	}

	// メンバーを取得
	member, err := j.readOnlyRepository.FindMemberByID(input.MemberID)
	if err != nil {
		output.Err = err
		j.presenter.Execute(output)
		return output
	}

	//メンバー名をセット
	output.MemberName = member.Name()

	// ゲームマスターを取得
	gameMaster, err := j.readOnlyRepository.FindGameMasterByGroupID(input.GroupRoomID)
	if err != nil {
		output.Err = err
		j.presenter.Execute(output)
		return output
	}

	// 参加メンバーを追加
	gameMaster.SetMember(member.ID())
	if _, err = j.gameMasterRepository.Save(gameMaster); err != nil {
		output.Err = err
		j.presenter.Execute(output)
		return output
	}

	// 出力
	output.Err = nil
	j.presenter.Execute(output)
	return output
}
