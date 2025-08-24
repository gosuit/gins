package gins

import (
	"github.com/gin-gonic/gin"
	"github.com/gosuit/e"
)

// Abort handles errors by logging them and responding with an appropriate JSON error message.
// It checks the error code and logs a specific message based on the type of error.
func Abort(c *gin.Context, err e.Error) {
	if err.GetStatus() == e.Internal {
		GetL(c).Error("Something going wrong", err.SlErr())
	} else {
		GetL(c).Info("Invalid input data", err.SlErr())
	}

	c.AbortWithStatusJSON(
		err.GetHttpCode(),
		err.ToJson(),
	)
}
