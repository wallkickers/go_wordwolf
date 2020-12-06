package infrastruture

import (
	"github.com/go-server-dev/src/app/interface_adapter"
	"net/http"
)

// Router URLに対するルーティング
type Router struct {
	controller interface_adapter.LinebotController
}

// NewRouter コンストラクタ
func addLineBotController(c interface_adapter.LinebotController) *Router {
	return &Router{controller: c}
}

// Init ルーティング設定
func (r *Router) Init() {
	http.HandleFunc("/callback", r.controller.CallBack)
}
