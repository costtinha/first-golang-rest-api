package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})

			}
		}()
		ctx.Next()
	}
}
