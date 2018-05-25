package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oujia/gobase"
	"github.com/oujia/gobase-demo/config"
)

func main()  {
	r := gin.Default()

	r.Use(config.DwLog.NewSelfLog())

	gobase.InitRouter(r, routers)

	r.Run(config.SERVER_ADDR)
}