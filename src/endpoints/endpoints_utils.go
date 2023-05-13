package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func badRequestIfError(err *error, context *gin.Context) {
	if *err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": (*err).Error()})
		return
	}
}

func exceptionRecover() {
	err := recover()

	// Log package to be applied
	fmt.Println(err)
}
