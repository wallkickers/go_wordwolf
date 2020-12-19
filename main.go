package main

import (
	"fmt"
	"github.com/go-server-dev/src/app/domain"
	"github.com/go-server-dev/src/app/infrastructure"
	"github.com/go-server-dev/src/app/interface_adapter"
	"github.com/go-server-dev/src/app/mocks"
	"github.com/go-server-dev/src/app/usecase/join"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"os"
	"strconv"
)

// httpリクエストが来た時
func handler(w http.ResponseWriter, r *http.Request) {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	logrus.Fatalf("Error loading env: %v", err)
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	logrus.Fatalf("Error loading env: %v", err)
	// }
	bot, err := linebot.New(
		// TODO: 開発者のLINEアカウントごとに変更する必要あり
		os.Getenv("LINEBOT_CHANNEL_SECRET"),
		os.Getenv("LINEBOT_CHANNEL_ACCESS_TOKEN"),
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
				replyMessage := "メッセージID：" + message.ID + "メッセージ：" + message.Text
				// エラー処理
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

// func main() {
// 	port, _ := strconv.Atoi(os.Args[1])
// 	fmt.Printf("Starting server at Port %d", port)
// 	http.HandleFunc("/", handler)             // / にリクエストが来た時はhandlerを呼ぶ。→ Hello world
// 	http.HandleFunc("/callback", lineHandler) // /callbackにリクエストが来た時にlineHandlerを呼ぶ
// 	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
// }

func main() {
	// ダミーデータ
	dummyMember := domain.NewMember("1", "テスト太郎")
	dummyGameMaster := domain.NewGameMaster("123", domain.GroupRoomType("group"))
	// 参照用リポジトリMock
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	readOnlyRepositoryMock.On("FindMemberByID", mock.AnythingOfType("int")).Return(dummyMember, nil)
	readOnlyRepositoryMock.On("FindGameMasterByGroupID", mock.AnythingOfType("int")).Return(dummyGameMaster, nil)
	// 更新用リポジトリMock
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	gameMasterRepositoryMock.On("Save", dummyGameMaster).Return(nil)
	// Presenter
	joinPresenter := interface_adapter.NewLineBotJoinPresenter()
	// UseCase
	joinUseCase := join.NewUseCaseImpl(gameMasterRepositoryMock, readOnlyRepositoryMock, joinPresenter)
	// Controller
	controller := interface_adapter.NewLinebotController(joinUseCase)
	// Router
	router := infrastructure.Router{}
	router.AddLineBotController(*controller)
	router.Init()

	port, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Starting server at Port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
