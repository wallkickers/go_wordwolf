// Package join ゲーム参加に関するパッケージ
package join

// Presenter ゲーム参加ユースケース表示
type Presenter interface {
	// 参加状況を表示
	Execute(Output) error
}
