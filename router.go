package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/oujia/gobase-demo/controller"
	"github.com/oujia/gobase"
	"github.com/oujia/gobase-demo/config"
)

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
		"gobase/remote",
		func(c *gin.Context) {

			dwHttp := &gobase.DwHttp{c, config.DwLog}
			resp, err := dwHttp.SimpleGet("http://www.duowan.com")
			if err != nil {
				gobase.NewResponseWithMSG(gobase.CODE_REQUEST_TIMEOUT, nil, err.Error()).SendBy(c)
			} else {
				gobase.NewResponse(gobase.CODE_SUCCESS, resp).SendBy(c)
			}
		},
	},
	gobase.Router{
		http.MethodGet,
		"gobase/hdPrize",
		controller.HDPrize,
	},
	gobase.Router{
		http.MethodGet,
		"gobase/hdPrizeOne",
		controller.HDPrizeOne,
	},
	gobase.Router{
		http.MethodGet,
		"loc/episodeTest",
		controller.EpisodeTest,
	},
	gobase.Router{
		http.MethodGet,
		"loc/testCount",
		controller.TestGetCount,
	},
	gobase.Router{
		http.MethodGet,
		"loc/testUpdate",
		controller.TestUpdate,
	},
	gobase.Router{
		http.MethodGet,
		"loc/testDel",
		controller.TestDel,
	},
	gobase.Router{
		http.MethodGet,
		"loc/testAdd",
		controller.TestAdd,
	},
	gobase.Router{
		http.MethodGet,
		"loc/testAdds",
		controller.TestAdds,
	},
	gobase.Router{
		http.MethodGet,
		"r2m/testGetRow",
		controller.TestR2MGetRow,
	},
}
