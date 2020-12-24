// Package accept_votes ゲーム参加に関するパッケージ
package accept_votes

// Presenter ゲーム参加ユースケース表示
type Presenter interface {
	// 参加状況を表示
	Execute(Output) error
}
