package controller

import (
	"github.com/oujia/gobase-demo/model"
	"github.com/oujia/gobase-demo/config"
	"github.com/gin-gonic/gin"
	"github.com/oujia/gobase"
)

var episodeR2m *model.EpisodeR2m

func init()  {
	r2mInfo, ok := config.GlobalConf.R2mInfo["mikan-episode"]
	if !ok {
		panic("lost mikan-episode r2minfo")
	}
	episodeR2m = new(model.EpisodeR2m)
	episodeR2m.TableHelper = new(gobase.TableHelper)
	episodeR2m.Redis = config.REDIS_LOCAL
	episodeR2m.R2mInfo = r2mInfo
	episodeR2m.DB = config.DB_LOCALHOST
	episodeR2m.TableName = "episode"
	episodeR2m.DbKey = "mikan"
}

func TestR2MGetRow(c *gin.Context) {
	var hash string
	where := map[string]interface{} {
		"id": 205979,
	}
	keyword := map[string]interface{}{
		"_field": "hash",
	}
	err := episodeR2m.R2M.GetRow(&hash, where, keyword)
	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	gobase.NewResponse(gobase.CODE_SUCCESS, hash).SendBy(c)
}
