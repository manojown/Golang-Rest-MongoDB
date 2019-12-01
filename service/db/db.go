package database

import (
	"github.com/globalsign/mgo"
)

// var database string = ""
// var db *mgo.Database

func NewDB(url string, database string) (*mgo.Database, error) {
	db, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	c := db.DB(database)
	return c, nil
}

// func Collection(CollectionName string) *mgo.Collection {

// 	fmt.Println(CollectionName)
// 	return db.C(CollectionName)

// }
