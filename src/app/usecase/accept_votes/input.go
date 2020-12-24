// Package accept_votes ゲーム参加に関するパッケージ
package accept_votes

// Input 投票するユースケース入力DTO
type Input struct {
	ReplyToken    string
	FromMemberID  string
	ToMemberID    string
	GroupRoomID   string
	GroupRoomType string
}
