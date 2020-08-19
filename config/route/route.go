package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mint-leaf/gin-tpl/lib"
)

var router *gin.Engine
var factory *MiddlewareFactory

func init() {
	router = gin.Default()
	// middleware
	factory = new(MiddlewareFactory)
	router.Use(
		factory.AllowOption(),
	)
}

// GetBaseRoute get router source
func GetBaseRoute() *gin.RouterGroup {
	return router.Group("/api")
}

func GetApp() *gin.Engine {
	return router
}

// WrapHandler transform handler function results into standard structure
func WrapHandler(runner func(ctx *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if data, err := runner(c); err == nil {
			if data == nil {
				data = map[string]interface{}{}
			}
			c.JSON(http.StatusOK, map[string]interface{}{
				"data":    data,
				"message": "OK",
			})
		} else {
			e, ok := err.(*gin.Error)
			if !ok {
				e = lib.NewServiceError(err).(*gin.Error)
			}
			// default err is service error
			// and just bind status 400 for param error simply
			// can add other situations
			// todo: function that can provide code and error map
			code := http.StatusBadGateway
			switch e.Type {
			case gin.ErrorTypeBind:
				code = http.StatusBadRequest
			}
			c.AbortWithStatusJSON(code, e)
		}
	}
}
