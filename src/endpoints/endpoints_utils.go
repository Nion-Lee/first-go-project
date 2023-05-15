package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetBadRequestWithError(err error, context *gin.Context) {
	context.JSON(http.StatusBadRequest, gin.H{"error": (err).Error()})
}

func exceptionRecover() {
	err := recover()

	// Log package to be applied
	fmt.Println(err)
}
