package routers

import (
	"first-go-project/src/endpoints"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", endpoints.HelloWorld)
	router.GET("/Customer/:uuid", endpoints.GetCustomer)
	router.PUT("/Customer/Create", endpoints.CreateCustomer)
	router.PATCH("/Customer/Update", endpoints.UpdateCustomer)
	router.DELETE("/Customer/Delete/:uuid", endpoints.DeleteCustomer)
}
