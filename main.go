package main

import (
	"fmt"
	"log"
	"os"
	"restApi/handler"
	database "restApi/service/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	db, err := database.NewDB(os.Getenv("MONGODB"), os.Getenv("DATABASE"))

	if err != nil {
		fmt.Printf("errr ::: ", err.Error())
		log.Panic(err)
	}

	// env := &Env{db: db}
	r := gin.Default()
	r.GET("/user", handler.GetUser(db))
	r.POST("/user", handler.CreateUser(db))
	r.PUT("/user", handler.UpdateUser())
	r.DELETE("/user", handler.DeleteUser())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
