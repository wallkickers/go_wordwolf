// Package join ゲーム参加に関するパッケージ
package join

// UseCase ゲーム参加ユースケース
type UseCase interface {
	// ゲームに参加する
	Excute(Input) Output
}
