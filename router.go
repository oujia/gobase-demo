package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/oujia/gobase-demo/controller"
	"github.com/oujia/gobase"
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

			dwHttp := &gobase.DwHttp{c, DwLog}
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
}
