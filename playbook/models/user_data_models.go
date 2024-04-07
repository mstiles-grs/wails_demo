package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UsersInfo struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	SQLID int `json:"sql_id" bson:"sql_id"`
	UserName string `json:"user_name" bson:"user_name"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName string `json:"last_name" bson:"last_name"`

}