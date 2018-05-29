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
	ids := make([]int, 0)
	keyword := map[string]interface{}{
		"_foundRows": true,
		"_field": "id",
		"_sort": "id desc",
		"_limit": 5,
		"_where": "id>200000",
	}
	err := episodeDao.GetCol(&ids, nil, keyword)
	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	total := episodeDao.GetFoundRows()

	var hash string
	where := map[string]interface{} {
		"id": 205979,
	}
	keyword = map[string]interface{}{
		"_field": "hash",
	}
	err = episodeDao.GetOne(&hash, where, keyword)
	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	gobase.NewResponse(gobase.CODE_SUCCESS, map[string]interface{} {
		"idList": ids,
		"hashOne": hash,
		"total": total,
	}).SendBy(c)
}

func TestGetCount(c *gin.Context)  {
	where := map[string]interface{}{
		"id": []int32{206152, 206147},
		"hash": []string{"c29d7563df79acee322b6da43a5a779211e4ae91", "0ac5f01e26f587d95430154b0944f6fd4f4ab2ac"},
	}
	keyword := map[string]interface{}{
		"_foundRows": true,
		"_field": "id",
		"_sort": "id desc",
		"_limit": 5,
		"_where": "id>200000",
	}
	count, err := episodeDao.GetCount(where, keyword)

	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	gobase.NewResponse(gobase.CODE_SUCCESS, count).SendBy(c)
}

func TestUpdate(c *gin.Context) {
	where := map[string]interface{}{
		"id": 532,
	}
	data := map[string]interface{} {
		"bangumiId": 123,
		"title": "test",
	}

	ar, err := episodeDao.UpdateObject(data, where)
	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	gobase.NewResponse(gobase.CODE_SUCCESS, ar).SendBy(c)
}


func TestDel(c *gin.Context) {
	where := map[string]interface{}{
		"id": 532,
	}

	ar, err := episodeDao.DelObject(where)
	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	gobase.NewResponse(gobase.CODE_SUCCESS, ar).SendBy(c)
}

func TestAdd(c *gin.Context)  {
	episode := model.Episode{
		Hash: "foobar1",
		Title: "test title1",
	}

	id, err := episodeDao.AddObject(episode)
	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	gobase.NewResponse(gobase.CODE_SUCCESS, id).SendBy(c)
}