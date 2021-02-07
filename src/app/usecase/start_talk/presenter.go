package start_talk

// Presenter 出力インターフェース
type Presenter interface {
	Execute(Output)
	FinishTalk(FinishTalkOutput)
}
