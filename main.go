package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"mathb-go/router"
)

func main() {
	engine := gin.Default()
	router.RegisterRouters(engine)
	lo.Must0(engine.Run(":7156"))
}
