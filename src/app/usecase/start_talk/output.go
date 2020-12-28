package start_talk

import "time"

// Output ユースケース出力DTO
type Output struct {
	ReplyToken    string
	MemberID      string
	GroupRoomID   string
	GroupRoomType string
	TalkTimeMin   time.Duration
	Err           error
}
