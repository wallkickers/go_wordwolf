// Package join ゲーム参加に関するパッケージ
package join

import (
	"github.com/go-server-dev/src/app/domain"
	"github.com/go-server-dev/src/app/usecase/repository"
)

// UseCaseImpl ゲーム参加ユースケース実装
type UseCaseImpl struct {
	gameMasterRepository repository.GameMasterRepository
	memberRepository     repository.MemberRepository
	joinRepository       ReadOnlyRepository
	presenter            Presenter
}

// インターフェースを満たしているかのチェック
var _ UseCase = (*UseCaseImpl)(nil)

// NewUseCaseImpl ゲーム参加ユースケースインスタンスを新規作成する
func NewUseCaseImpl(
	gameMasterRepository repository.GameMasterRepository,
	memberRepository repository.MemberRepository,
	joinRepository ReadOnlyRepository,
	presenter Presenter) *UseCaseImpl {
	return &UseCaseImpl{
		gameMasterRepository: gameMasterRepository,
		memberRepository:     memberRepository,
		joinRepository:       joinRepository,
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
	member, errMemeberRepo := j.joinRepository.FindMemberByID(input.MemberID)
	if errMemeberRepo != nil {
		if errMemeberRepo == ErrNotFound {
			// 未登録時、新規作成する
			member = domain.NewMember(input.MemberID)
		} else {
			output.err = errMemeberRepo
			j.presenter.Execute(output)
			return output
		}
	}

	// ゲームマスターを取得
	gameMaster, errGameMasterRepo := j.joinRepository.FindGameMasterByGroupID(input.GroupRoomID)
	if errGameMasterRepo != nil {
		if errGameMasterRepo == ErrNotFound {
			// 未登録時、新規作成する
			gameMaster = domain.NewGameMaster(input.GroupRoomID, domain.GroupRoomType(input.GroupRoomType))
		} else {
			// 取得エラー時の処理
			output.err = errGameMasterRepo
			j.presenter.Execute(output)
			return output
		}
	}

	// 参加メンバーに追加
	gameMaster.SetMember(member.ID())

	// 保存
	j.gameMasterRepository.Save(gameMaster)

	// 出力
	output.err = nil
	j.presenter.Execute(output)
	return output
}
