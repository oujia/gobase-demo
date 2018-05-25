package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/oujia/gobase"
	"github.com/oujia/gobase-demo/model"
	"github.com/oujia/gobase-demo/config"
)

var prizeDao *model.PrizeDao

func init()  {
	prizeDao = new(model.PrizeDao)
	prizeDao.TableHelper.DB = config.DB_HD
	prizeDao.TableName = "prize"
}

func HDPrize(c *gin.Context)  {
	list := []model.Prize{}

	err := prizeDao.GetSome(&list)
	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	gobase.NewResponse(gobase.CODE_SUCCESS, map[string]interface{}{
		"list": list,
	}).SendBy(c)

}

func HDPrizeOne(c *gin.Context)  {
	prize := model.Prize{}

	err := prizeDao.GetOne(&prize)
	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	gobase.NewResponse(gobase.CODE_SUCCESS, prize).SendBy(c)
}
