package apis

import (
	"net/http"

	"github.com/ajay-kumar-ks/gin-test/models"
	"github.com/gin-gonic/gin"
)

func listUser(ctx *gin.Context) {	
		user1 := models.User{
		"ajay",
		"ajay@gmail.com",
	}
		user2 := models.User{
		"ajay",
		"ajay@gmail.com",
	}
	ctx.JSON(http.StatusOK, []models.User{user1,user2})
}
func createUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
