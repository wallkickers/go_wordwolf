package interface_adapter

import (
	"fmt"

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
func (p *LineBotStartTalkPresenter) Execute(output startTalk.Output) error {

	if output.Err != nil {
		switch output.Err {
		// 既に参加済み
		case startTalk.ErrAlreadyStarted:
			replyMessage := "既にトークはスタートしています。"
			if _, err := p.bot.ReplyMessage(output.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
				return err
			}
			return startTalk.ErrAlreadyStarted
		default:
			return output.Err
		}
	}

	replyMessage := "トークを開始します。\r\n制限時間は" + fmt.Sprintf("%1.0f", output.TalkTimeMin.Minutes()) + "分です。"
	if _, err := p.bot.ReplyMessage(output.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		return err
	}
	return nil
}
