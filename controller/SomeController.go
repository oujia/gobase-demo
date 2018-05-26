package controller

import (
	"github.com/oujia/gobase-demo/model"
	"github.com/oujia/gobase-demo/config"
	"github.com/gin-gonic/gin"
	"github.com/oujia/gobase"
)

var episodeDao *model.EpisodeDao

func init()  {
	episodeDao = new(model.EpisodeDao)
	episodeDao.TableHelper.DB = config.DB_LOCALHOST
	episodeDao.TableName = "episode"
}

func EpisodeTest(c *gin.Context)  {
	var id int
	keyword := map[string]interface{}{
		"_field": "id",
		"_sort": "id desc",
		"_limit": 5,
	}
	err := episodeDao.GetOne(&id, nil, keyword)
	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	gobase.NewResponse(gobase.CODE_SUCCESS, id).SendBy(c)
}