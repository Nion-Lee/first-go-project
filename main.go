package main

import (
	"First_Go_Project/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.RegisterRoutes(router)

	router.Run(":8080")
}
