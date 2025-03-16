package gin

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/lec"
	"github.com/gosuit/sl"
)

const (
	CtxKey = "lec-ctx"
)

func InitLogger(c lec.Context) gin.HandlerFunc {
	log := c.Logger()

	log.Info("logger middleware enabled.")

	return func(c *gin.Context) {
		ctx := lec.New(log)

		c.Set(CtxKey, ctx)

		req := c.Request

		c.Next()
		entry := log.With(
			sl.StringAttr("method", req.Method),
			sl.StringAttr("path", req.URL.Path),
			sl.StringAttr("remote_addr", req.RemoteAddr),
			sl.StringAttr("user_agent", req.UserAgent()),
		)

		t1 := time.Now()
		defer func() {
			entry.Info("request completed",
				sl.IntAttr("status", c.Writer.Status()),
				sl.StringAttr("duration", time.Since(t1).String()),
			)
		}()
	}
}

func GetCtx(c *gin.Context) lec.Context {
	if c, ok := c.Get(CtxKey); ok {
		return c.(lec.Context)
	}

	return lec.New(sl.Default())
}

func GetL(c *gin.Context) sl.Logger {
	ctx := GetCtx(c)

	return ctx.Logger()
}
