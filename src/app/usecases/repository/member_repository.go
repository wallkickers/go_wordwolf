package repository

import (
	"github.com/takuyaaaaaaahaaaaaa/go_wordwolf/src/app/domain"
)

// MemberRepository メンバーの更新リポジトリ（取得には使わないこと）
type MemberRepository interface {
	// メンバーを保存する
	Save(domain.Member) error
	// メンバーIDを指定してメンバーを取得する
	FindMemberByID(string) (domain.Member, error)
}
