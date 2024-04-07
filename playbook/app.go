package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mstiles-grs/wails_demo/playbook/mongoDB"
	"github.com/mstiles-grs/wails_demo/playbook/sqlDB"
	"go.mongodb.org/mongo-driver/mongo"
)

  type App struct {
    ctx context.Context
  }

      var sqlInstance *sql.DB
      var mongoClient *mongo.Client
      var ctx context.Context
      var cancel context.CancelFunc
      var database string

      func NewApp() *App {
        return &App{}
      }

  // startup is called when the app starts. The context is saved
  // so we can call the runtime methods
  func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
    sqlInstance = sqlDB.SQLStartUp()
    foundInstance, dbName, mongoErr := mongoDB.MongoStartUp()

    if mongoErr != nil {
      if mongoErr.Error() == "Failed To Connect To Mongo" {
        fmt.Println("Failed To connect")
      } else if mongoErr.Error() == "Faild To Ping Mongo" {
        fmt.Println("Failed To ping")
      } else {
        fmt.Println("Failed")
      }
    }

    mongoClient = foundInstance
    database = dbName

    UserCollection := mongoClient.Database(database).Collection("User")

    mongoUser, findUserErr := mongoDB.GetMongoUser(ctx, mongoClient, UserCollection)

    if findUserErr != nil {
      fmt.Println("Can't find Users")
    }

    fmt.Println(mongoUser)
  }

  func (a *App) shutdown(ctx context.Context) {
    mongoDB.DisconnectMongoDB(mongoClient)
  }


  // Greet returns a greeting for the given name
  func (a *App) Greet(name string) string {
    return fmt.Sprintf("Hello %s, It's show time!", name)
}
