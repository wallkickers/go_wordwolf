package repository

import (
	// "time"
	"github.com/go-server-dev/src/app/domain"
)

// GameMasterRepository ゲームマスタ リポジトリ
type GameMasterRepository interface {
	// ゲームマスターを保存する
	Save(*domain.GameMaster)(bool, error)
	// StartToMesureTime(*domain.GameMaster) error
	// GetLimitTime(*domain.GameMaster) (time.Time, error)
}
