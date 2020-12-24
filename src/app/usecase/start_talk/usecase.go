// Package start_talk トーク開始に関するパッケージ
package start_talk

// UseCase ゲーム参加ユースケース
type UseCase interface {
	Excute(Input) Output
}
