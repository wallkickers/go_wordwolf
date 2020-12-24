package start_talk

// Output ユースケース出力DTO
type Output struct {
	ReplyToken    string
	MemberID      string
	GroupRoomID   string
	GroupRoomType string
	Err           error
}
