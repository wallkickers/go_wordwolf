package domain

/*
状態管理

SettingTime（設定フェーズ）
↓
TalkTime（トークフェーズ）
↓↑
VoteTime（投票フェーズ）
↓
ResultTime（結果フェーズ）
↓
*/

// State ゲーム状態
type State string

const (
	settingTime = State("settingTime")
	talkTime    = State("talkTime")
	voteTime    = State("voteTime")
	resultTime  = State("resultTime")
)

// StartTalk トーク開始
func (g *GameMaster) StartTalk() {
	g.state = talkTime
}

// EndTalk トーク終了
func (g *GameMaster) EndTalk() {
	g.state = voteTime
}

// IsTalkTime トークタイム
func (g *GameMaster) IsTalkTime() bool {
	if g.state == talkTime {
		return true
	}
	return false
}
