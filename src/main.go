package main

import (
	"first-go-project/src/repositories"
	"first-go-project/src/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	repositories.InitializeDB()

	router := gin.Default()
	routers.RegisterRoutes(router)

	router.Run(":8080")

	fmt.Println("Done")
}
