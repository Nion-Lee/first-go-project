package endpoints

import (
	"First_Go_Project/dtos"
	"First_Go_Project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(context *gin.Context) {
	defer exceptionRecover()

	var dto dtos.CustomerDTO
	err := context.BindJSON(&dto)
	badRequestIfError(&err, context)

	newDTO, err := services.CreateCustomer(&dto)
	badRequestIfError(&err, context)

	context.JSON(http.StatusOK, newDTO)
}
