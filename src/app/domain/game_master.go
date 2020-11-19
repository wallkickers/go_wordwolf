// Package domain ビジネスルールに関するパッケージ
package domain

import (
	"time"
)

// GameMaster はゲーム管理者を表す
type GameMaster struct {
	// グループID
	groupID string
	// 参加メンバーリスト
	memberList map[string]bool
	// 各メンバーのお題リスト([参加メンバー]お題)
	themeManagement map[string]string
	// 各メンバーの投票先([投票元]投票先)
	voteManagement map[string]string
	// ゲームプレイ時間
	talkTimeMin time.Time
	// ゲーム終了時刻
	endTime time.Time
}

// NewGameMaster ゲーム管理者インスタンスを新規作成する
func NewGameMaster(groupID string) *GameMaster {
	return &GameMaster{
		groupID: groupID,
	}
}

// GroupID グループIDを取得する
func (g *GameMaster) GroupID() string {
	return g.groupID
}

// MemberList 参加メンバーリストを取得する
func (g *GameMaster) MemberList() map[string]bool {
	return g.memberList
}

// SetMember 参加メンバーを設定する
func (g *GameMaster) SetMember(memberID string) {
	g.memberList[memberID] = true
}

// DeleteMember 離脱メンバーを設定する
func (g *GameMaster) DeleteMember(memberID string) {
	delete(g.memberList, memberID)
}

// ThemeManagement 各メンバーのお題リストを取得する
func (g *GameMaster) ThemeManagement() map[string]string {
	return g.themeManagement
}

// SetThemeManagement 各メンバーのお題リストを設定する
func (g *GameMaster) SetThemeManagement(themeManagement map[string]string) {
	g.themeManagement = themeManagement
}

// VoteManagement 各メンバーの投票先を取得する([投票元]投票先)
func (g *GameMaster) VoteManagement() map[string]string {
	return g.voteManagement
}

// SetVoteManagement 各メンバーの投票先を設定する
func (g *GameMaster) SetVoteManagement(fromMemberID string, toMemberID string) {
	g.voteManagement[fromMemberID] = toMemberID
}

// TalkTimeMin ゲームプレイ時間(分)を取得する
func (g *GameMaster) TalkTimeMin() time.Time {
	return g.talkTimeMin
}

// SetTalkTimeMin ゲームプレイ時間(分)を設定する
func (g *GameMaster) SetTalkTimeMin(talkTimeMin time.Time) {
	g.talkTimeMin = talkTimeMin
}

// EndTime ゲーム終了時刻を取得する
func (g *GameMaster) EndTime() time.Time {
	return g.endTime
}

// SetEndTime ゲーム終了時刻を設定する
func (g *GameMaster) SetEndTime(endTime time.Time) {
	g.endTime = endTime
}

// RemainingTime ゲーム残り時間を取得する
func (g *GameMaster) RemainingTime() time.Duration {
	return g.endTime.Sub(time.Now())
}
