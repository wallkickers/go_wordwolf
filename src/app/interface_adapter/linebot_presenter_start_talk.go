package interface_adapter

import (
	"fmt"
	"log"

	startTalk "github.com/go-server-dev/src/app/usecase/start_talk"
	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBotStartTalkPresenter LINEBOTプレゼンター
type LineBotStartTalkPresenter struct {
	bot *linebot.Client
}

// インターフェースを満たしているかのチェック
var _ startTalk.Presenter = (*LineBotStartTalkPresenter)(nil)

// NewLineBotStartTalkPresenter  コンストラクタ
func NewLineBotStartTalkPresenter(bot *linebot.Client) *LineBotStartTalkPresenter {
	return &LineBotStartTalkPresenter{
		bot: bot,
	}
}

// Execute 表示処理
func (p *LineBotStartTalkPresenter) Execute(output startTalk.Output) {

	if output.Err != nil {
		switch output.Err {
		// 既に参加済み
		case startTalk.ErrAlreadyStarted:
			replyMessage := "既にトークはスタートしています。"
			log.Print("ID：" + output.GroupRoomID + "token：" + output.ReplyToken + "：" + replyMessage)
			if _, err := p.bot.ReplyMessage(output.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
				print(err)
				return
			}
			print(output.Err)
			return
		default:
			print(output.Err)
			return
		}
	}

	replyMessage := "トークを開始します。\r\n制限時間は" + fmt.Sprintf("%1.0f", output.TalkTimeMin.Minutes()) + "分です。"
	log.Print("ID：" + output.GroupRoomID + " token：" + output.ReplyToken + "：" + replyMessage)
	if _, err := p.bot.PushMessage(output.GroupRoomID, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Print(err)
		return
	}
	return
}

// FinishTalk 終了通知処理
func (p *LineBotStartTalkPresenter) FinishTalk(output startTalk.FinishTalkOutput) {

	replyMessage := "トーク時間が終了しました。"
	log.Print("ID：" + output.GroupRoomID + "token：" + output.ReplyToken + "：" + replyMessage)
	if _, err := p.bot.PushMessage(output.GroupRoomID, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Print(err)
		return
	}
	return
}
