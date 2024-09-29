package main

import (
	"fmt"

	// "github.com/ajay-kumar-ks/gin-test/apis"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("skill map")
	router := gin.Default()

	apis.ResgiterUserApis(router)

	router.Run()
}