package route

import "github.com/gin-gonic/gin"

type MiddlewareFactory struct{}

// AllowOption allow option method
func (factory *MiddlewareFactory) AllowOption() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
		}
	}
}
