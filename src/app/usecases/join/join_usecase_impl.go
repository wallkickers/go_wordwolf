// Package join ゲーム参加に関するパッケージ
package join

import (
	"github.com/takuyaaaaaaahaaaaaa/go_wordwolf/src/app/domain"
)

type JoinUseCaseImpl struct {
	gameMasterRepository GameMasterRepository
}

// ゲームに参加する
func (u *JoinUseCaseImpl) joinGame(JoinInput) JoinOutput {
	// ゲームマスターを取得する
	// ゲームマスターにメンバーを追加する
	// メンバーを保存する
	// ゲームマスタを保存する。
	// プレゼンターに出力する
}
