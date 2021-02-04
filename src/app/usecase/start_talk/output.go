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

// FinishTalkOutput トーク終了出力DTO
type FinishTalkOutput struct {
	ReplyToken    string
	GroupRoomID   string
	GroupRoomType string
	Err           error
}
