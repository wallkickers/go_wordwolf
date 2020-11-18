package join_usecase

import (
	"github.com/takuyaaaaaaahaaaaaa/go_wordwolf/src/app/domain"
)

type JoinUseCaseImpl struct {
	gameMasterRepository GameMasterRepository
}

// ゲームに参加する
func (u *JoinUseCaseImpl) joinGame(JoinInputDto) JoinOutputDto {
	// ゲームマスター排他制御
	// ゲームマスターを取得する
	// ゲームマスターにメンバーを追加する
	// ゲームマスターの状態を保存する。
	// プレゼンターに保存する
}
