package repository

import (
	"github.com/takuyaaaaaaahaaaaaa/go_wordwolf/src/app/domain"
)

// MemberRepository メンバーの更新リポジトリ（取得には使わないこと）
type MemberRepository interface {
	// メンバーを保存する
	save(domain.Member) error
	// メンバーIDを指定してメンバーを取得する
	findMemberByID(string) (domain.Member, error)
}
