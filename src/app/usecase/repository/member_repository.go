package repository

import (
	"github.com/go-server-dev/src/app/domain"
)

// MemberRepository メンバーリポジトリ
type MemberRepository interface {
	// メンバーを保存する
	Save(*domain.Member) error
	// メンバーIDを指定してメンバーを取得する
	FindMemberByID(string) (*domain.Member, error)
}
