package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-server-dev/src/app/domain"
	"github.com/go-server-dev/src/app/infrastructure"
	"github.com/go-server-dev/src/app/infrastructure/database"
	"github.com/go-server-dev/src/app/interface_adapter"
    "github.com/go-server-dev/src/app/mocks"
	// "github.com/go-server-dev/src/app/usecase/repository"
	accept_votes "github.com/go-server-dev/src/app/usecase/accept_votes/impl"
	join "github.com/go-server-dev/src/app/usecase/join/impl"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/stretchr/testify/mock"

	// mongodbテスト用import
	"context" // manage multiple requests
	// "reflect" // get an object type
    // "time"
    // "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: 開発用：httpリクエストが来た時
func handler(w http.ResponseWriter, r *http.Request) {
	// .envから環境変数読み込み
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading env: %v", err)
	}
	fmt.Fprintf(w, "Hello world\n")
}

func newLineBot() *linebot.Client {
	// .envから環境変数読み込み
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading env: %v", err)
	}
	bot, err := linebot.New(
		os.Getenv("LINEBOT_CHANNEL_SECRET"),
		os.Getenv("LINEBOT_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
	// .envから環境変数読み込み
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading env: %v", err)
	}
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
	// LINEBotをNew
    linebot := newLineBot()
    
    // DB Connect
	db, err := infrastructure.Connect()
	if err != nil {
		logrus.Infof("Error connecting DB: %v", err)
		// Heroku用 アプリの起動に合わせてDBが起動できないことがあるので再接続を試みる
		// db, _ = infrastructure.Connect()
    }
    defer db.Disconnect(context.Background())

	// ゲーム参加UseCase
	dummyMember := domain.NewMember("1", "テスト太郎")
	dummyGameMaster := domain.NewGameMaster("123", domain.GroupRoomType("group"))
	readOnlyRepositoryMock := new(mocks.ReadOnlyRepository)
	readOnlyRepositoryMock.On("FindMemberByID", mock.AnythingOfType("string")).Return(dummyMember, nil)
    readOnlyRepositoryMock.On("FindGameMasterByGroupID", mock.AnythingOfType("string")).Return(dummyGameMaster, nil)
	gameMasterRepositoryMock := new(mocks.GameMasterRepository)
	gameMasterRepositoryMock.On("Save", dummyGameMaster).Return(nil)
	gameMasterRepositoryMock.On("GetLimitTime", dummyGameMaster).Return(nil)
	gameMasterRepositoryMock.On("StartToMesureTime", dummyGameMaster).Return(nil)

    joinPresenter := interface_adapter.NewLineBotJoinPresenter(linebot)
    gameMasterRepository := database.NewGameMasterRepository(db)

	// joinUseCase := join.NewUseCaseImpl(gameMasterRepositoryMock, readOnlyRepositoryMock, joinPresenter)
	joinUseCase := join.NewUseCaseImpl(gameMasterRepository, readOnlyRepositoryMock, joinPresenter)

    // 投票受付UseCase
    // gameMasterRepository := new(repository.GameMasterRepository)
    // gameMasterRepository := database.NewUserRepository(db)
	acceptVotesPresenter := interface_adapter.NewLineBotAcceptVotesPresenter(linebot)
	acceptVotesUseCase := accept_votes.NewUseCaseImpl(gameMasterRepository, readOnlyRepositoryMock, acceptVotesPresenter)

	// Controller
	controller := interface_adapter.NewLinebotController(
		acceptVotesUseCase,
		joinUseCase,
		linebot,
    )

	// // DB Connect
	// db, err := infrastructure.Connect()
	// if err != nil {
	// 	logrus.Infof("Error connecting DB: %v", err)
	// 	// Heroku用 アプリの起動に合わせてDBが起動できないことがあるので再接続を試みる
	// 	// db, _ = infrastructure.Connect()
    // }
    // defer db.Disconnect(context.Background())
	// output sql query
    // db.LogMode(true)

	// Connect to the MongoDB and return Client instance
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// fmt.Println("mongo.Connect() ERROR:", err)
	// os.Exit(1)
	// }

	// Declare Context type object for managing multiple API requests
	// ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// // Access a MongoDB collection through a database
	// col := client.Database("wordwolf").Collection("game_master")
    // fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	// InsertOne() method Returns mongo.InsertOneResult
	// result, insertErr := col.InsertOne(ctx, oneDoc)
	// result, insertErr := col.InsertOne(ctx, bson.D{
    //     {"title", "The Polyglot Developer Podcast"},
    //     {"author", "Nic Raboy"},
    //     {"tags", bson.A{"development", "programming", "coding"}},
    // })
	// if insertErr != nil {
    //     fmt.Println("InsertOne ERROR:", insertErr)
	//     os.Exit(1) // safely exit script on error
	// } else {
	//     fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
	//     fmt.Println("InsertOne() API result:", result)
	
	// // get the inserted ID string
	// newID := result.InsertedID
	// fmt.Println("InsertOne() newID:", newID)
	// fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))

	// Router 
	router := infrastructure.Router{}
	// router := Initialize(db)
	router.AddLineBotController(*controller)
	router.Init()

	port, _ := strconv.Atoi(os.Args[1])
	http.HandleFunc("/", handler)             // / にリクエストが来た時はhandlerを呼ぶ。→ Hello world
	fmt.Printf("Starting server at Port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
