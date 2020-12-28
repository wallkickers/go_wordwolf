// Package domain ビジネスルールに関するパッケージ
package domain

import (
	"time"
)

// GameMaster はゲーム管理者を表す
type GameMaster struct {
	// グループ・ルームID
	groupRoomID string
	// グループ・ルーム種別
	groupRoomType GroupRoomType
	// 参加メンバーリスト
	memberList map[string]bool
	// 各メンバーのお題リスト([参加メンバー]お題)
	themeManagement map[string]string
	// 各メンバーの投票先([投票元]投票先)
	voteManagement map[string]string
	// ゲームプレイ時間
	talkTimeMin time.Duration
	// 状態
	state State
}

// GroupRoomType グループ・ルームの種別
type GroupRoomType string

const (
	//Group 種別がグループ（）
	group = GroupRoomType("group")
	//Room 種別がルーム（複数人のトーク）
	room = GroupRoomType("room")
)

// NewGameMaster ゲーム管理者インスタンスを新規作成する
func NewGameMaster(groupRoomID string, groupRoomType GroupRoomType) *GameMaster {
	return &GameMaster{
		groupRoomID:     groupRoomID,
		groupRoomType:   groupRoomType,
		memberList:      map[string]bool{},
		themeManagement: map[string]string{},
		voteManagement:  map[string]string{},
		talkTimeMin:     time.Minute * 3, //3分間
	}
}

// GroupRoomID グループルームIDを取得する
func (g *GameMaster) GroupRoomID() string {
	return g.groupRoomID
}

// GroupRoomType グループルーム種別を取得する
func (g *GameMaster) GroupRoomType() GroupRoomType {
	return g.groupRoomType
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
func (g *GameMaster) TalkTimeMin() time.Duration {
	return g.talkTimeMin
}

// SetTalkTimeMin ゲームプレイ時間(分)を設定する
func (g *GameMaster) SetTalkTimeMin(talkTimeMin time.Duration) {
	g.talkTimeMin = talkTimeMin
}
