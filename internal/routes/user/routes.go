package user

import (
	"api/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.Handle("POST", "/")
}

func GetUser(c *gin.Context) {
	userId, exists := c.GetQuery("user_id")
	if !exists {
		c.JSON(utils.ValidationError("User id is required", nil))
	}
	
	user := GetUser
	
}

func CreateUser(c *gin.Context) {

}
