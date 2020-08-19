package ping

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mint-leaf/gin-tpl/config/route"
)

//all router in package register here
func init() {
	group := route.GetBaseRoute().Group("/test")
	group.GET("/ping/:name", route.WrapHandler(testGetRoute))

}

func testGetRoute(ctx *gin.Context) (interface{}, error) {
	return fmt.Sprintf("hello %s", ctx.Query("name")), nil
}
