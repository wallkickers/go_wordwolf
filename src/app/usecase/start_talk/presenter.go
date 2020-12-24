// Package join ゲーム参加に関するパッケージ
package join

// Presenter 出力インターフェース
type Presenter interface {
	// 参加状況を表示
	Execute(Output) error
}
