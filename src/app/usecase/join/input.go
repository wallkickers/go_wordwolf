// Package join ゲーム参加に関するパッケージ
package join

// Input ゲーム参加ユースケース入力DTO
type Input struct {
	ReplyToken    string
	MemberID      string
	GroupRoomID   string
	GroupRoomType string
}
