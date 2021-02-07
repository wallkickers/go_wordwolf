package domain

/*
状態管理

PrepareTime（設定時間）
↓
TalkTime（トーク時間）
↓↑
VoteTime（投票時間）
↓
ResultTime（結果時間）
↓
*/

// State ゲーム状態
type State string

const (
	prepareTime = State("prepareTime")
	talkTime    = State("talkTime")
	voteTime    = State("voteTime")
	resultTime  = State("tesultTime")
)

// StartTalk トーク開始
func (g *GameMaster) StartTalk() {
	g.state = talkTime
}

// IsTalkTime トークタイム
func (g *GameMaster) IsTalkTime() bool {
	if g.state == talkTime {
		return true
	}
	return false
}
