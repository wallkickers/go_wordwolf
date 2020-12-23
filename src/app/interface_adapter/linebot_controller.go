package interface_adapter

import (
	"log"
	"net/http"

	"github.com/go-server-dev/src/app/usecase/accept_votes"
	"github.com/go-server-dev/src/app/usecase/join"
	"github.com/line/line-bot-sdk-go/linebot"
)

// LinebotController LINEBOTコントローラ
type LinebotController struct {
	acceptVotesUseCase accept_votes.UseCase // 投票を受け付ける
	joinUseCase join.UseCase // ゲームに参加する
	bot         *linebot.Client
}

// NewLinebotController コンストラクタ
func NewLinebotController(
	acceptVotesUseCase accept_votes.UseCase,
	joinUseCase join.UseCase,
	bot *linebot.Client,
	) *LinebotController {
	return &LinebotController{
		acceptVotesUseCase: acceptVotesUseCase,
		joinUseCase: joinUseCase,
		bot:         bot,
	}
}

// CallBack LINEBOTに関するリクエスト処理
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
	c.handleEvent(events)
}

// handleEvent Eventを振り分ける
func (c *LinebotController) handleEvent(events []*linebot.Event) {
	for _, event := range events {
		log.Printf("Got event %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			//メッセージイベント
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				c.handleText(message, event.ReplyToken, event.Source)
			//その他イベント
			default:
				log.Printf("Unknown message: %v", message)
			}
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}

// handleText テキストメッセージを処理する
func (c *LinebotController) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) {
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
			GroupRoomType: string(source.Type),
		}
		c.joinUseCase.Excute(input)
	case "投票":
		input := accept_votes.Input{
			ReplyToken:    replyToken,
			FromMemberID:  source.UserID,
			ToMemberID:    source.UserID,
			GroupRoomID:   groupRoomID,
			GroupRoomType: string(source.Type),
		}
		c.acceptVotesUseCase.Excute(input)
	case "テスト":
		replyMessage := "テストメッセージ：" + message.Text
		if _, err := c.bot.ReplyMessage(replyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
			log.Print(err)
		}
	default:
	}
}
