package interfaces

import (
	"github.com/go-server-dev/src/app/usecase/join"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

// LinebotController LINEBOTコントローラ
type LinebotController struct {
	joinUseCase join.UseCase // ゲームに参加する
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
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := app.handleText(message, event.ReplyToken, event.Source); err != nil {
					log.Print(err)
				}
			default:
				log.Printf("Unknown message: %v", message)
			}
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}

func (c *LinebotController) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	var groupRoomID string
	switch source.Type {
	case "group":
		groupRoomID = source.GroupID
	case "room":
		groupRoomID = source.RoomID
	}

	switch message.Text {
	case "参加":
		input := join.Input{
			ReplyToken:    replyToken,
			MemberID:      source.UserID,
			GroupRoomID:   groupRoomID,
			GroupRoomType: linebot.EventSourceType(source.Type)}
	default:
	}
}
