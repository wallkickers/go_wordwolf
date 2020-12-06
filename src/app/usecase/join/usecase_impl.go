// Package join ゲーム参加に関するパッケージ
package join

import (
	"github.com/go-server-dev/src/app/usecase/repository"
)

// UseCaseImpl ゲーム参加ユースケース実装
type UseCaseImpl struct {
	gameMasterRepository repository.GameMasterRepository
	readOnlyRepository   ReadOnlyRepository
	presenter            Presenter
}

// インターフェースを満たしているかのチェック
var _ UseCase = (*UseCaseImpl)(nil)

// NewUseCaseImpl ゲーム参加ユースケースインスタンスを新規作成する
func NewUseCaseImpl(
	gameMasterRepository repository.GameMasterRepository,
	readOnlyRepository ReadOnlyRepository,
	presenter Presenter) *UseCaseImpl {
	return &UseCaseImpl{
		gameMasterRepository: gameMasterRepository,
		readOnlyRepository:   readOnlyRepository,
		presenter:            presenter,
	}
}

// Excute ゲームに参加する
func (j *UseCaseImpl) Excute(input Input) Output {

	//出力DTO作成
	output := Output{
		MemberID:      input.MemberID,
		GroupRoomID:   input.GroupRoomID,
		GroupRoomType: input.GroupRoomType,
		ReplyToken:    input.ReplyToken,
	}

	// メンバーを取得
	member, err := j.readOnlyRepository.FindMemberByID(input.MemberID)
	if err != nil {
		output.err = err
		j.presenter.Execute(output)
		return output
	}

	//メンバー名をセット
	output.MemberName = member.Name()

	// ゲームマスターを取得
	gameMaster, err := j.readOnlyRepository.FindGameMasterByGroupID(input.GroupRoomID)
	if err != nil {
		output.err = err
		j.presenter.Execute(output)
		return output
	}

	// 参加メンバーを追加
	gameMaster.SetMember(member.ID())
	if err = j.gameMasterRepository.Save(gameMaster); err != nil {
		output.err = err
		j.presenter.Execute(output)
		return output
	}

	// 出力
	output.err = nil
	j.presenter.Execute(output)
	return output
}
