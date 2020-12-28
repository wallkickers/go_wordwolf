// Package accept_votes 投票するパッケージ
package accept_votes

// UseCase 投票するユースケース
type UseCase interface {
	// 投票する
	Excute(Input) Output
}
