package interface_adapter

import (
	"github.com/go-server-dev/src/app/usecase/join"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

// LineBotJoinPresenter LINEBOTプレゼンター
type LineBotJoinPresenter struct {
	bot *linebot.Client
}

// インターフェースを満たしているかのチェック
var _ join.Presenter = (*LineBotJoinPresenter)(nil)

// NewLineBotJoinPresenter  コンストラクタ
func NewLineBotJoinPresenter() *LineBotJoinPresenter {

	secret := os.Getenv("LBOT_SECRET")
	token := os.Getenv("LBOT_TOKEN")
	bot, err := linebot.New(secret, token)

	if err != nil {
		// TODO:環境設定読み込み失敗時のエラー処理
		log.Fatal(err)
	}

	return &LineBotJoinPresenter{
		bot: bot,
	}
}

// Execute 表示処理
func (p *LineBotJoinPresenter) Execute(output join.Output) error {

	if output.Err != nil {
		switch output.Err {
		// 既に参加済み
		case join.ErrIsExisted:
			replyMessage := output.MemberName + "さんは既に参加済みです。"
			if _, err := p.bot.ReplyMessage(output.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
				return err
			}
			return join.ErrIsExisted
		default:
			return output.Err
		}
	}

	replyMessage := output.MemberName + "さんの参加を受け付けました。"
	if _, err := p.bot.ReplyMessage(output.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		return err
	}
	return nil
}
