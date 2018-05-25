package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oujia/gobase"
)

func HDPrize(c *gin.Context)  {
	prizeDao := new(PrizeDao)
	prizeDao.TableHelper.DB = DB_HD
	prizeDao.TableName = "prize"
	list := []Prize{}

	err := prizeDao.GetSome(&list)
	if err != nil {
		gobase.NewResponse(gobase.CODE_DB_ERROR, err.Error()).SendBy(c)
		return
	}

	gobase.NewResponse(gobase.CODE_SUCCESS, map[string]interface{}{
		"list": list,
	}).SendBy(c)

}
