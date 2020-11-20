// Package join ゲーム参加に関するパッケージ
package join

import (
	"github.com/go-server-dev/src/app/domain"
	"github.com/go-server-dev/src/app/usecases/repository"
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

// Excute ゲームに参加する
func (j *UseCaseImpl) Excute(input Input) Output {

	//出力DTO作成
	var output Output
	output.MemberName = input.MemberName
	output.MemberID = input.MemberID
	output.GroupID = input.GroupID
	output.ReplyToken = input.ReplyToken

	// メンバーを取得
	member, errMemeberRepo := j.memberRepository.FindMemberByID(input.MemberID)
	if errMemeberRepo != nil {
		if errMemeberRepo == repository.ErrNotFound {
			// 未登録時、新規作成する
			member = domain.NewMember(input.MemberID, input.MemberName)
		} else {
			// 取得エラー時の処理
			// TODO エラーをセット output.err = ???
			j.presenter.Execute(output)
			return output
		}
	}

	// ゲームマスターを取得
	gameMaster, errGameMasterRepo := j.gameMasterRepository.FindGameMasterByGroupID(input.GroupID)
	if errGameMasterRepo != nil {
		if errGameMasterRepo == repository.ErrNotFound {
			// 未登録時、新規作成する
			gameMaster = domain.NewGameMaster(input.GroupID)
		} else {
			// 取得エラー時の処理
			// TODO エラーをセット output.err = ???
			j.presenter.Execute(output)
			return output
		}
	}

	// メンバー名を更新
	member.SetName(input.MemberName)
	// 参加メンバーに追加
	gameMaster.SetMember(member.ID())

	// 保存
	j.memberRepository.Save(member)
	j.gameMasterRepository.Save(gameMaster)

	// 出力
	output.err = nil
	j.presenter.Execute(output)
	return output
}
