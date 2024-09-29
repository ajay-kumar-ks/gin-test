package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ResgiterUserApis(router *gin.Engine){
	userGroup := router.Group("api/users/")
	fmt.Println(userGroup)
	userGroup.GET("",listUser)
	userGroup.POST("",createUser)
}