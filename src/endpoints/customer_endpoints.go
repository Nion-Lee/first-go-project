package endpoints

import (
	"first-go-project/src/dtos"
	"first-go-project/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCustomer(context *gin.Context) {
	defer exceptionRecover()

	uuidStr := context.Param("uuid")
	_, err := uuid.Parse(uuidStr)
	if err != nil {
		SetBadRequestWithError(err, context)
		return
	}

	dto, err := services.GetCustomer(&uuidStr)
	if err != nil {
		SetBadRequestWithError(err, context)
		return
	}

	context.JSON(http.StatusOK, dto)
}

func CreateCustomer(context *gin.Context) {
	defer exceptionRecover()

	var dto dtos.CustomerDTO
	err := context.BindJSON(&dto)
	if err != nil {
		SetBadRequestWithError(err, context)
		return
	}

	uid, err := services.CreateCustomer(&dto)
	if err != nil {
		SetBadRequestWithError(err, context)
		return
	}

	context.JSON(http.StatusOK, uid)
}

// 測試用
func HelloWorld(context *gin.Context) {
	response := services.HelloWorld()
	context.JSON(http.StatusOK, response)
}
