package user

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.Handle("POST", "/")
}

func GetUser() {

}

func CreateUser(c *gin.Context) {

}
