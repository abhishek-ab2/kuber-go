package main

import (
	"api/internal/db"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)


func main(){
	ctx := context.Background()

	db := db.GetDB()

	if err := db.Connect(ctx, "api"); err!=nil{
		log.Fatal("Failed to connect to DB awwww")
	}
	defer db.Disconnect(ctx)

	app := gin.Default()

	app.Run(":5000")


}
