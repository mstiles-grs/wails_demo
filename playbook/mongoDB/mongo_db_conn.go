package mongoDB

import (
	"context"
	"fmt"
	"log"

	"github.com/mstiles-grs/wails_demo/playbook/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoStartUp() (*mongo.Client, string, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)
if err != nil {
	return nil, "", fmt.Errorf("Failed To Connect To Mongo")
}

// Check the connection
err = client.Ping(context.TODO(), nil)
if err != nil {
	return nil, "", fmt.Errorf("Failed To Ping Mongo")
}

return client, "wails_test", nil
}


func DisconnectMongoDB(mongoClient *mongo.Client) {
    if err := mongoClient.Disconnect(context.TODO()); err != nil {
        log.Fatalf("Error disconnecting from MongoDB: %v", err)
    }
}

func GetMongoUser(ctx context.Context, mongoClient *mongo.Client, userCollection *mongo.Collection) ([]models.UsersInfo, error) {

	var users []models.UsersInfo

	filter := bson.M{}


	cur, err := userCollection.Find(ctx, filter)

	if err != nil {

		return nil, err

	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {

		var user models.UsersInfo

		if err := cur.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil

}
