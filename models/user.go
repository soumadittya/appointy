package models

import "gopkg.in/mgo.v2.bson"

type User struct{
	Id bson.objectId `json:"id" bson:"_id"`
	Name string      `json:"name" bson:"name"`
	Email string    `json:"email" bson:"email"`
	Password int          `json:"password" bson:"password"`
}