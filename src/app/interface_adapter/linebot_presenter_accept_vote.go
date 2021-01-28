package interface_adapter

import (
	"github.com/go-server-dev/src/app/usecase/accept_votes"
	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBotAcceptVotesPresenter LINEBOTプレゼンター
type LineBotAcceptVotesPresenter struct {
	bot *linebot.Client
}

// インターフェースを満たしているかのチェック
var _ accept_votes.Presenter = (*LineBotAcceptVotesPresenter)(nil)

// NewLineBotAcceptVotesPresenter  コンストラクタ
func NewLineBotAcceptVotesPresenter(bot *linebot.Client) *LineBotAcceptVotesPresenter {
	return &LineBotAcceptVotesPresenter{
		bot: bot,
	}
}

// Execute 表示処理
func (p *LineBotAcceptVotesPresenter) Execute(output accept_votes.Output) error {

	if output.Err != nil {
		switch output.Err {
		// 既に投票済み
		case accept_votes.ErrIsExisted:
			replyMessage := "既に投票済みです。"
			if _, err := p.bot.ReplyMessage(output.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
				return err
			}
			return accept_votes.ErrIsExisted
		default:
			return output.Err
		}
	}

	replyMessage := "投票を受け付けました。"
	if _, err := p.bot.ReplyMessage(output.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		return err
	}
	return nil
}
