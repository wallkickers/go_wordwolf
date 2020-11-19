// Package join ゲーム参加に関するパッケージ
package join

import (
	"github.com/takuyaaaaaaahaaaaaa/go_wordwolf/src/app/domain"
	"github.com/takuyaaaaaaahaaaaaa/go_wordwolf/src/app/usecases/repository"
)

type JoinUseCaseImpl struct {
	member               domain.Member
	gameMaster           domain.GameMaster
	gameMasterRepository repository.GameMasterRepository
	memberRepository     repository.MemberRepository
}

// インターフェースを満たしているかのチェック
var _ JoinUseCase = (*JoinUseCaseImpl)(nil)

// ゲームに参加する
func (j *JoinUseCaseImpl) JoinGame(input JoinInput) JoinOutput {
	// ゲームマスターを取得する
	var gameMaster, err = j.gameMasterRepository.FindGameMasterByGroupID(input.GroupID)
	if err != nil {
		if err == repository.ErrNotFound

	}
	if gameMaster == nil {
		gameMaster = domain.NewGameMaster(input.GroupID)
	}

	// メンバーを取得する
	var ga, err = j.gameMasterRepository.FindGameMasterByGroupID(input.GroupID)
	if err != nil {
	}
	if gameMaster == nil {
		gameMaster = domain.NewGameMaster(input.GroupID)
	}
	// ゲームマスターにメンバーを追加する
	// メンバーを保存する
	// ゲームマスタを保存する。
	// プレゼンターに出力する
	var output = new(JoinOutput)
	return *output
}
