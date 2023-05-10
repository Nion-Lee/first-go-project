package routers

import (
	"first-go-project/endpoints"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", endpoints.HelloWorld)
	router.PUT("/Customer/Create", endpoints.CreateCustomer)

}
