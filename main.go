package main

import (
	"first-go-project/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.RegisterRoutes(router)

	router.Run(":8080")
}
