// Package join ゲーム参加に関するパッケージ
package join

import (
	"github.com/go-server-dev/src/app/domain"
	"github.com/go-server-dev/src/app/usecase/repository"
)

// UseCaseImpl ゲーム参加ユースケース実装
type UseCaseImpl struct {
	member               domain.Member
	gameMaster           domain.GameMaster
	gameMasterRepository repository.GameMasterRepository
	memberRepository     repository.MemberRepository
	presenter            Presenter
}

// インターフェースを満たしているかのチェック
var _ UseCase = (*UseCaseImpl)(nil)

// NewUseCaseImpl ゲーム参加ユースケースインスタンスを新規作成する
func NewUseCaseImpl(member domain.Member,
	gameMaster domain.GameMaster,
	gameMasterRepository repository.GameMasterRepository,
	memberRepository repository.MemberRepository,
	presenter Presenter) *UseCaseImpl {
	return &UseCaseImpl{
		member:               member,
		gameMaster:           gameMaster,
		gameMasterRepository: gameMasterRepository,
		memberRepository:     memberRepository,
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
	member, errMemeberRepo := j.memberRepository.FindMemberByID(input.MemberID)
	if errMemeberRepo != nil {
		if errMemeberRepo == repository.ErrNotFound {
			// 未登録時、新規作成する
			member = domain.NewMember(input.MemberID)
		} else {
			// 取得エラー時の処理
			// TODO: エラーをセット output.err = ???
			j.presenter.Execute(output)
			return output
		}
	}
	// TODO: ユーザ名を取得して、更新

	// ゲームマスターを取得
	gameMaster, errGameMasterRepo := j.gameMasterRepository.FindGameMasterByGroupID(input.GroupRoomID)
	if errGameMasterRepo != nil {
		if errGameMasterRepo == repository.ErrNotFound {
			// 未登録時、新規作成する
			gameMaster = domain.NewGameMaster(input.GroupRoomID, domain.GroupRoomType(input.GroupRoomType))
		} else {
			// 取得エラー時の処理
			// TODO エラーをセット output.err = ???
			j.presenter.Execute(output)
			return output
		}
	}

	// 参加メンバーに追加
	gameMaster.SetMember(member.ID())

	// 保存
	j.memberRepository.Save(member)
	j.gameMasterRepository.Save(gameMaster)

	// 出力
	//output.memberName = ?? 取得したユーザ名を入力
	output.err = nil
	j.presenter.Execute(output)
	return output
}
