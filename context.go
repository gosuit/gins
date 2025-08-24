package gins

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/c"
	"github.com/gosuit/sl"
)

const (
	// Key used to store the c.Context in the gin.Context
	CtxKey = "ctx"
)

// InitLogger initializes a logger middleware for Gin.
// It logs the request details and duration after the request is completed.
// It also set c.Context to gin.Context.
func InitLogger(ctx c.Context) gin.HandlerFunc {
	log := ctx.Logger()

	log.Info("logger middleware enabled.")

	return func(ginCtx *gin.Context) {
		ginCtx.Set(CtxKey, c.New(log))

		req := ginCtx.Request

		ginCtx.Next()
		entry := log.With(
			sl.StringAttr("method", req.Method),
			sl.StringAttr("path", req.URL.Path),
			sl.StringAttr("remote_addr", req.RemoteAddr),
			sl.StringAttr("user_agent", req.UserAgent()),
		)

		t1 := time.Now()
		defer func() {
			entry.Info("request completed",
				sl.IntAttr("status", ginCtx.Writer.Status()),
				sl.StringAttr("duration", time.Since(t1).String()),
			)
		}()
	}
}

// GetCtx retrieves the c.Context from the Gin context.
// If no context is found, it returns a new with default Logger.
func GetCtx(ginCtx *gin.Context) c.Context {
	if ctx, ok := ginCtx.Get(CtxKey); ok {
		return ctx.(c.Context)
	}

	return c.New(sl.Default())
}

// GetL retrieves the logger from the c.Context stored in the Gin context.
func GetL(ginCtx *gin.Context) sl.Logger {
	return GetCtx(ginCtx).Logger()
}
