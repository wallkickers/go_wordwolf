package repository

import (
	"github.com/takuyaaaaaaahaaaaaa/go_wordwolf/src/app/domain"
)

// GameMasterRepository ゲームマスターの更新リポジトリ（取得には使わないこと）
type GameMasterRepository interface {
	// ゲームマスターを保存する
	save(domain.GameMaster) error
	// グループIDを指定してゲームマスターを取得する
	findGameMasterByGroupID(string) (domain.GameMaster, error)
}
