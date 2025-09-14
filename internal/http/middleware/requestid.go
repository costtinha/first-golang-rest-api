package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const HeaderRequestId = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetHeader(HeaderRequestId)
		if id == "" {
			id = uuid.NewString()
		}
		ctx.Writer.Header().Set(HeaderRequestId, id)
		ctx.Set(HeaderRequestId, id)
		ctx.Next()
	}
}
