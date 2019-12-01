package handler

import (
	"log"
	"restApi/contract"
	"restApi/service/db/models"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type User contract.User

func GetUser(db *mgo.Database) gin.HandlerFunc {

	return func(c *gin.Context) {

		Response := contract.Response{}
		query := bson.M{}

		result, err := models.GetUser(db, query)
		Response.Message = "All user get successfully."
		Response.Data = result
		if err != nil {
			log.Panic("Error while called Getuser Model")
		}
		// fmt.Printf("result is %v", result)
		c.JSON(200, Response)

	}

}

func CreateUser(db *mgo.Database) gin.HandlerFunc {

	return func(c *gin.Context) {
		Response := contract.Response{}
		user := bson.M{}
		c.Bind(&user)
		// fmt.Print("query is:: ", user)
		result, err := models.CreateUser(db, user)
		Response.Message = result
		if err != nil {
			log.Panic("Error while called Getuser Model")
		}
		// fmt.Printf("result is %v", result)
		c.JSON(200, Response)
	}

}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get data save",
		})
	}

}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get data save",
		})
	}

}
