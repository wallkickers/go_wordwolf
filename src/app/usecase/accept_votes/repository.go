package accept_votes

import (
	"github.com/go-server-dev/src/app/domain"
)

// ReadOnlyRepository 参照用リポジトリ
type ReadOnlyRepository interface {
	// グループIDを指定してゲームマスターを取得する
	FindGameMasterByGroupID(string) (*domain.GameMaster, error)
}
