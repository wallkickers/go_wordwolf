package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "github.com/line/line-bot-sdk-go/linebot"
)

// httpリクエストが来た時
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!\n")
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
    bot, err := linebot.New(
        "ae5b824deb86c7caff7f9c3a8098fb9c",
        "6PK3v2YKfi+bAxCV7scReyKDvR0YRyBONhbjLSPOEnTc52VcrpMYvItJu6/yWF+UKwY8/DVQ/OBNfV7cs/fc57Bj9IojD1818zyte6dtrXDc1JV/BQeJmJX+mAVlCf1WqKqx6QMeZLPad90vPndRiAdB04t89/1O/w1cDnyilFU=",
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
            // case *linebot.StickerMessage:
            //     replyMessage := fmt.Sprintf(
            //         // ↓vindになる。%xにカンマ区切りで記載した順で入る。
            //         "sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType
            //     )

            //     // エラー処理
            //     if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
            //         log.Print(err)
            //     }
            }
        }
    }
}

func main() {
    port, _ := strconv.Atoi(os.Args[1])
    fmt.Printf("Starting server at Port %d", port)
    http.HandleFunc("/", handler) // / にリクエストが来た時はhandlerを呼ぶ。→ Hellop world
    http.HandleFunc("/callback", lineHandler) // /callbackにリクエストが来た時にlineHandlerを呼ぶ
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}