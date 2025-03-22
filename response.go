package gins

import (
	"github.com/gin-gonic/gin"
	"github.com/gosuit/e"
)

// Abort handles errors by logging them and responding with an appropriate JSON error message.
// It checks the error code and logs a specific message based on the type of error.
func Abort(c *gin.Context, err e.Error) {
	log := GetL(c)

	if err.GetCode() == e.Internal {
		err.Log("Something going wrong...")
	} else {
		log.Info("Invalid input data")
	}

	c.AbortWithStatusJSON(
		err.ToHttpCode(),
		err.ToJson(),
	)
}
