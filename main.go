package main

import (
	"first-go-project/repositories"
	"first-go-project/routers"
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
