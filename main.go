package main

import (
	"fmt"

	// "github.com/ajay-kumar-ks/gin-test/apis"
	"github.com/ajay-kumar-ks/gin-test/database"
	"github.com/ajay-kumar-ks/gin-test/handlers"
	"github.com/ajay-kumar-ks/gin-test/managesrs"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("skill map")

	database.Initialise()
	router := gin.Default()

	userManger := managesrs.NewUserManager()
	userHandler := handlers.NewUserHanlderFrom(userManger)
	userHandler.RegisterUserapis(router)
	// apis.ResgiterUserApis(router)

	router.Run()
}