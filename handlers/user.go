package handlers

import (
	"fmt"
	"net/http"

	"github.com/ajay-kumar-ks/gin-test/common"
	// "github.com/ajay-kumar-ks/gin-test/database"
	"github.com/ajay-kumar-ks/gin-test/managesrs"
	// "github.com/ajay-kumar-ks/gin-test/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName string
	userManager *managesrs.UserManager
}

func NewUserHanlderFrom(userManager *managesrs.UserManager) *UserHandler{
	return &UserHandler{
		"api/users",
	    userManager,
	}
	
}

func (userHandler *UserHandler) RegisterUserapis(r *gin.Engine){
	userGroup := r.Group(userHandler.groupName)
	userGroup.POST("", userHandler.Create)
}
func (userHandler *UserHandler) Create(ctx *gin.Context){

	userData := common.NewUserCreationInput()

	err := ctx.BindJSON(&userData)

	if err!=nil{
		fmt.Println("failed to bind data")
	}
	
	newUser, err := userHandler.userManager.Create(userData)

	if err!=nil{
		fmt.Println("failed to create user")
	}

	// database.DB.Create(&models.User{FullName: userData.FullName, Email: userData.Email})
	ctx.JSON(http.StatusOK,newUser)
}