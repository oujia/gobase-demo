package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oujia/gobase"
	"net/http"
	"github.com/jmoiron/sqlx"
)

var DwLog *gobase.DwLog
var DB_LOCALHOST *sqlx.DB
var GlobalConf *gobase.GlobalConf

func init()  {
	conf, err := gobase.LoadGlobalConf(GLOBAL_CONFIG_PATH)
	if err != nil {
		panic(err)
	}
	GlobalConf = conf

	DB_LOCALHOST, err = gobase.NewDbClient("localhost", GlobalConf.DbInfo)
	if err != nil {
		panic(err)
	}
}

var routers = gobase.Routers{
	gobase.Router{
		http.MethodGet,
		"/gobase/get",
		func(c *gin.Context) {
			gobase.NewResponse(gobase.CODE_NORMAL_ERROR, gin.H{
				"test": 123,
				"foo": "bar",
				"call_id": c.Request.Header.Get("call_id"),
			}).SendBy(c)
		},
	},
	gobase.Router{
		http.MethodPost,
		"gobase/post",
		func(c *gin.Context) {
			gobase.NewResponseWithMSG(gobase.CODE_SUCCESS, nil,"测试成功").SendBy(c)
		},
	},
	gobase.Router{
		http.MethodGet,
		"gobase/db",
		func(c *gin.Context) {

			publishGroupDao := new(PublishGroupDao)
			publishGroupDao.TableHelper.DB = DB_LOCALHOST
			publishGroupDao.TableName = "publishGroup"
			pgList := []PublishGroup{}
			err := publishGroupDao.GetAll(&pgList, "", 10)

			episodeDao := new(EpisodeDao)
			episodeDao.TableHelper.DB = DB_LOCALHOST
			episodeDao.TableName = "episode"
			episodeList := []Episode{}

			err = episodeDao.GetAll(&episodeList, "id, hash, title", 5)
			if err != nil {
				gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
				return
			}

			gobase.NewResponse(gobase.CODE_SUCCESS, map[string]interface{}{
				"publishGroup": pgList,
				"episode": episodeList,
			}).SendBy(c)
		},
	},
	gobase.Router{
		http.MethodGet,
		"gobase/remote",
		func(c *gin.Context) {

			dwHttp := &gobase.DwHttp{c, DwLog}
			resp, err := dwHttp.SimpleGet("http://www.duowan.com")
			if err != nil {
				gobase.NewResponseWithMSG(gobase.CODE_REQUEST_TIMEOUT, nil, err.Error()).SendBy(c)
			} else {
				gobase.NewResponse(gobase.CODE_SUCCESS, resp).SendBy(c)
			}
		},
	},
}

func main()  {
	r := gin.Default()

	logRedis, ok := GlobalConf.RedisInfo["logstash_redis"]
	if !ok {
		panic("lost log redis config")
	}

	DwLog = &gobase.DwLog{
		LogKey: LOG_KEY,
		SelfCall: LOG_SELF_CALL,
		ModuleCall: LOG_MODULE_CALL,
		RedisClient: gobase.NewRedisClient(&logRedis),
	}

	r.Use(DwLog.NewSelfLog())

	gobase.InitRouter(r, routers)

	r.Run(SERVER_ADDR)
}
