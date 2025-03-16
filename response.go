package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/gosuit/e"
)

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
