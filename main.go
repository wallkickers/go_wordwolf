package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
	"strconv"
)

// httpリクエストが来た時
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		// TODO: 開発者のLINEアカウントごとに変更する必要あり
		"2340b865b7e61c194f0ea792f7a46e68",
		"hnkIH+9nEkwhBJKGWf6OWbSZsv/y3j3ylruAwnb3QSb0ll8nl6zyshAcAnOv+biOzztBNQ0jx8fbXB18MqghDL3eaAC+qQ0CaNs2MQy4HTawotB76gyQv+WvLc4KC+spdQefReuw+UOMxBovgVZkLAdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	// range: foreachみたいなやつ。 _がindex eventがvalueになる。
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				replyMessage := message.Text
				// エラー処理
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Starting server at Port %d", port)
	http.HandleFunc("/", handler)             // / にリクエストが来た時はhandlerを呼ぶ。→ Hello world
	http.HandleFunc("/callback", lineHandler) // /callbackにリクエストが来た時にlineHandlerを呼ぶ
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}