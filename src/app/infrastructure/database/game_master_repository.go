package database

import (
	"fmt"
	"os"
	"time"
	"reflect" // get an object type

	"github.com/go-server-dev/src/app/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"context" // manage multiple requests
)



type GameMasterRepository struct {
	db *mongo.Client
}

func NewGameMasterRepository(db *mongo.Client) *GameMasterRepository {
	return &GameMasterRepository{db: db}
}

func (repo *GameMasterRepository) Save(gameMaster *domain.GameMaster) (bool, error) {
	// Access a MongoDB collection through a database
	client := repo.db
	col := client.Database("wordwolf").Collection("game_master")
	// fmt.Println("Collection type:", reflect.TypeOf(col), "\n")
	
	// Declare a MongoDB struct instance for the document's fields and data
	// oneDoc := MongoFields{
	// FieldStr: "SomeValue",
	// FieldInt: 12345,
	// FieldBool: true,
	// }
	// fmt.Println("oneDoc TYPE:", reflect.TypeOf(oneDoc), "\n")
	
	// InsertOne() method Returns mongo.InsertOneResult
	// result, insertErr := col.InsertOne(ctx, oneDoc)

	// type GameMaster struct {
	// 	// グループ・ルームID
	// 	groupRoomID string
	// 	// グループ・ルーム種別
	// 	groupRoomType GroupRoomType
	// 	// 参加メンバーリスト
	// 	memberList map[string]bool
	// 	// 各メンバーのお題リスト([参加メンバー]お題)
	// 	themeManagement map[string]string
	// 	// 各メンバーの投票先([投票元]投票先)
	// 	voteManagement map[string]string
	// 	// ゲームプレイ時間
	// 	talkTimeMin time.Time
	// 	// ゲーム終了時刻
	// 	endTime time.Time
	// }

	// Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	result, err := col.InsertOne(ctx, gameMaster)
	if err != nil {
		fmt.Println("InsertOne ERROR:", err)
		os.Exit(1) // safely exit script on error
		return false, err
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() API result:", result)
	
		// get the inserted ID string
		newID := result.InsertedID
		fmt.Println("InsertOne() newID:", newID)
		fmt.Println("InsertOne() newID type:", reflect.TypeOf(newID))

		return true, err
	}
}

// func (repo *GameMasterRepository) StartToMesureTime(gameMaster *domain.GameMaster) error {
// 	return
// }

// func (repo *GameMasterRepository) GetLimitTime(gameMaster *domain.GameMaster) (time.Time, error) {
// 	time := time.Time;
// 	err := nil
// 	return time, err
// }
