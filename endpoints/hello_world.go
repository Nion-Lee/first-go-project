package endpoints

import (
	"First_Go_Project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(context *gin.Context) {
	response := services.HelloWorld()
	context.JSON(http.StatusOK, response)
}
