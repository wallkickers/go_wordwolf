package line

import (
	"github.com/go-server-dev/src/app/usecase/join"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

// LinebotController LINEBOTコントローラ
type LinebotController struct {
	joinUseCase join.UseCase
	bot         *linebot.Client
}

// NewLinebotController コンストラクタ
func NewLinebotController(
	joinUseCase join.UseCase) *LinebotController {

	secret := os.Getenv("LBOT_SECRET")
	token := os.Getenv("LBOT_TOKEN")
	bot, err := linebot.New(secret, token)

	if err != nil {
		// TODO:環境設定読み込み失敗時のエラー処理
		log.Fatal(err)
	}

	return &LinebotController{
		joinUseCase: joinUseCase,
		bot:         bot,
	}
}

// CallBack LINEBOTに関する処理
func (c *LinebotController) CallBack(w http.ResponseWriter, r *http.Request) {
	events, err := c.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		log.Printf("Got event %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
		case linebot.EventTypeFollow:
		case linebot.EventTypeUnfollow:
		case linebot.EventTypeJoin:
		case linebot.EventTypeLeave:
		case linebot.EventTypePostback:
		case linebot.EventTypeBeacon:
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}
