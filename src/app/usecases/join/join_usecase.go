// Package join ゲーム参加に関するパッケージ
package join

// JoinUseCase ゲーム参加ユースケース
type JoinUseCase interface {
	// ゲームに参加する
	JoinGame(JoinInput) JoinOutput
}
