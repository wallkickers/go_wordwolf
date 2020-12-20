package infrastructure

import (
	"net/http"

	"github.com/go-server-dev/src/app/interface_adapter"
)

// Router URLに対するルーティング
type Router struct {
	controller interface_adapter.LinebotController
}

// AddLineBotController コンストラクタ
func (r *Router) AddLineBotController(controller interface_adapter.LinebotController) {
	r.controller = controller
}

// Init ルーティング設定
func (r *Router) Init() {
	http.HandleFunc("/callback", r.controller.CallBack)
}
