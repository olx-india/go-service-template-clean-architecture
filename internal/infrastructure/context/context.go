package context

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinContext struct {
	*gin.Context
}

func NewGinContext(c *gin.Context) (*GinContext, error) {
	return &GinContext{
		Context: c,
	}, nil
}

func (ctx *GinContext) JSON(code int, obj interface{}) {
	ctx.Context.JSON(code, obj)
}

func (ctx *GinContext) ShouldBindJSON(obj interface{}) error {
	return ctx.Context.ShouldBindJSON(obj)
}

func (ctx *GinContext) AbortWithStatus(code int) {
	ctx.Context.AbortWithStatus(code)
}

func (ctx *GinContext) Next() {
	ctx.Context.Next()
}

func (ctx *GinContext) Request() *http.Request {
	return ctx.Context.Request
}

func (ctx *GinContext) Writer() gin.ResponseWriter {
	return ctx.Context.Writer
}
