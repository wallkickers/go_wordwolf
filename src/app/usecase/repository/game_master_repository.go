package repository

import (
	"github.com/go-server-dev/src/app/domain"
)

// GameMasterRepository ゲームマスタ リポジトリ
type GameMasterRepository interface {
	// ゲームマスターを保存する
	Save(*domain.GameMaster) error
	// グループIDを指定してゲームマスターを取得する
	FindGameMasterByGroupID(string) (*domain.GameMaster, error)
}
