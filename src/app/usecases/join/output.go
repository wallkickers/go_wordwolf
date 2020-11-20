// Package join ゲーム参加に関するパッケージ
package join

// Output ゲーム参加ユースケース出力DTO
type Output struct {
	ReplyToken string
	MemberID   string
	MemberName string
	GroupID    string
	err        error
}
