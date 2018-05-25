package model

import (
	"github.com/oujia/gobase"
	"time"
)

type Prize struct {
	Id int `db:"id"`
	Name string `db:"name"`
	CreateTime time.Time `db:"create_time"`
}

type PrizeDao struct {
	gobase.TableHelper
}

func (dao *PrizeDao) GetSome(list *[]Prize) error {
	where := map[string]interface{}{
		"hd_id": 8,
	}
	keyword := map[string]interface{}{
		"_field": "id, name, create_time",
		"_limit": 5,
	}
	return dao.GetAll(list, where, keyword)
}

func (dao *PrizeDao) GetOne(list *Prize) error {
	where := map[string]interface{}{
		"hd_id": 8,
	}
	keyword := map[string]interface{}{
		"_field": "id, name, create_time",
		"_limit": 5,
	}
	return dao.GetRow(list, where, keyword)
}