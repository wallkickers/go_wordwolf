// Package domain ビジネスルールに関するパッケージ
package domain

// Member ゲーム参加者
type Member struct {
	// ID
	id string
	// ユーザ名
	name string
}

// Members ゲーム参加者集合
type Members []Member

// NewMember ゲーム参加者インスタンスを新規作成する
func NewMember(groupID, id, name string) Member {
	return Member{
		id:   id,
		name: name,
	}
}

// ID IDを取得する
func (m *Member) ID() string {
	return m.id
}

// Name ユーザ名を取得する
func (m *Member) Name() string {
	return m.name
}
