package user

import "github.com/gin-gonic/gin"


type UserHandler struct {
	userRepo
}

func RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.Handle("POST", "/",)
}


func GetUser(){
	var k string

}


func CreateUser(c *gin.Context){
	
}
