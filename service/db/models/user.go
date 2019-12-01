package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
	"restApi/contract"
)

type UserCollection interface {
	GetUser(db *mgo.Database, query bson.M) (Users, error)
	CreateUser(db *mgo.Database, query bson.M) (string, error)
}

var collection string = "users"

type Users []contract.User

func GetUser(db *mgo.Database, query bson.M) (Users, error) {
	var user Users
	collection := db.C(collection)
	err := collection.Find(query).All(&user)
	if err != nil {
		log.Panic("Something went wrong while insert to db")
	}
	return user, nil
}

func CreateUser(db *mgo.Database, query bson.M) (string, error) {

	collection := db.C(collection)
	err := collection.Insert(query)
	if err != nil {
		return "Failed while inserting data.", err
		// log.Panic("Something went wrong while insert to db")
	}
	return "User added successfully.", nil

}
