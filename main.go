package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oujia/gobase"
	"github.com/oujia/gobase-demo/config"
)

var DwLog *gobase.DwLog

func main()  {
	r := gin.Default()

	logRedis, ok := config.GlobalConf.RedisInfo["logstash_redis"]
	if !ok {
		panic("lost log redis config")
	}

	DwLog = &gobase.DwLog{
		LogKey: config.LOG_KEY,
		SelfCall: config.LOG_SELF_CALL,
		ModuleCall: config.LOG_MODULE_CALL,
		RedisClient: gobase.NewRedisClient(&logRedis),
	}

	r.Use(DwLog.NewSelfLog())

	gobase.InitRouter(r, routers)

	r.Run(config.SERVER_ADDR)
}