//+build wireinject

package main

import (
	"github.com/go-server-dev/src/app/infrastructure"
	"github.com/go-server-dev/src/app/infrastructure/database"
	"github.com/go-server-dev/src/app/usecase/repository"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var superSet = wire.NewSet(
	// Database
	database.NewGameMasterRepository,
	wire.Bind(new(repository.GameMasterRepository), new(*database.GameMasterRepository)),
)

// Initialize DI
func Initialize(db *mongo.Client) *infrastructure.Router {
	wire.Build(superSet)
	return &infrastructure.Router{}
}
