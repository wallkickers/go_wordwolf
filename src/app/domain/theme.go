// Package domain ビジネスルールに関するパッケージ
package domain

// Theme お題
type Theme struct {
	// 多数派のお題
	majorityTheme string
	// 少数派(ウルフ)のお題
	wolfTheme string
}

// NewTheme お題インスタンスを新規作成する
func NewTheme(majorityTheme, wolfTheme string) Theme {
	return Theme{
		majorityTheme: majorityTheme,
		wolfTheme:     wolfTheme,
	}
}

// MajorityTheme 多数派のお題を取得する
func (t *Theme) MajorityTheme() string {
	return t.majorityTheme
}

// WolfTheme 少数派(ウルフ)のお題を取得する
func (t *Theme) WolfTheme() string {
	return t.wolfTheme
}
