package routers

import (
	"First_Go_Project/endpoints"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", endpoints.HelloWorld)
	router.PUT("/Customer/Create", endpoints.CreateCustomer)

}

/*
	guid, err := uuid.NewRandom()
	ggininder := guid.String()
	parsed, err := uuid.Parse(ggininder)
*/
