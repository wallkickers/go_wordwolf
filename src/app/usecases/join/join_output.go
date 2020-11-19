// Package join ゲーム参加に関するパッケージ
package join

type JoinOutput struct {
	ReplyToken string
	MemberID   string
	MemberName string
	GroupID    string
	err        error
}
