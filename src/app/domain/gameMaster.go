// Package domain ビジネスルールに関するパッケージ
package domain

import (
	"time"
)

// GameMaster はゲーム管理者を表す
type GameMaster struct {
	// LINEグループID
	groupID int
	// 参加メンバーリスト
	memberList map[int]bool
	// 各メンバーのお題リスト
	themeManagement map[int]string
	// 各メンバーの投票先
	voteManagement map[int]int
	// ゲームプレイ時間
	talkTimeMin time.Time
	// ゲーム終了時刻
	endTime time.Time
}

// NewGameMaster GameMasterのインスタンスを新規作成する
func NewGameMaster(groupID int) *GameMaster {
	return &GameMaster{
		groupID: groupID,
	}
}

// GroupID グループIDを取得する
func (g *GameMaster) GroupID() int {
	return g.groupID
}

// SetGroupID グループIDを設定する
func (g *GameMaster) SetGroupID(groupID int) {
	g.groupID = groupID
}

// MemberList 参加メンバーリストを取得する
func (g *GameMaster) MemberList() map[int]bool {
	return g.memberList
}

// SetMember 参加メンバーを設定する
func (g *GameMaster) SetMember(member int) {
	g.memberList[member] = true
}

// DeleteMember 離脱メンバーを設定する
func (g *GameMaster) DeleteMember(member int) {
	delete(g.memberList, member)
}

// ThemeManagement 各メンバーのお題リストを取得する
func (g *GameMaster) ThemeManagement() map[int]string {
	return g.themeManagement
}

// SetThemeManagement 各メンバーのお題リストを設定する
func (g *GameMaster) SetThemeManagement(themeManagement map[int]string) {
	g.themeManagement = themeManagement
}

// VoteManagement 各メンバーの投票先を取得する
func (g *GameMaster) VoteManagement() map[int]int {
	return g.voteManagement
}

// SetVoteManagement 各メンバーの投票先を設定する
func (g *GameMaster) SetVoteManagement(fromMember int, toMember int) {
	g.voteManagement[fromMember] = toMember
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

// インターフェイス
type GameMasterAction interface {
	AssignTheme(groupID int) (string, error)
	StartToMesureTime(groupID int) error
	GetLimitTime(groupID int) (time.Time, error)
	GetResult(groupID int) error
	AddMember(id, groupID int) error
	//ManageVote(id, id int) error
}
