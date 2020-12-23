// Package accept_votes ゲーム参加に関するパッケージ
package accept_votes

// Output ゲーム参加ユースケース出力DTO
type Output struct {
	ReplyToken    string
	GroupRoomID   string
	GroupRoomType string
	Err           error
}
