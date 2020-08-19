package main

import (
	"github.com/mint-leaf/gin-tpl/config/route"

	// register router in main package
	_ "github.com/mint-leaf/gin-tpl/router/ping"
)

func main() {
	_ = route.GetApp().Run(":8080")
}
