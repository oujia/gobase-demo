package main

import (
	"github.com/oujia/gobase"
	"time"
)

type PublishGroup struct {
	Id int `db:"id"`
	Name string `db:"name"`
	CreateTime time.Time `db:"createTime"`
}

type PublishGroupDao struct {
	gobase.TableHelper
}

type Episode struct {
	Id int `db:"id"`
	Hash string `db:"hash"`
	Title string `db:"title"`
}

type EpisodeDao struct {
	gobase.TableHelper
}

type Prize struct {
	Id int `db:"id"`
	Name string `db:"name"`
	CreateTime time.Time `db:"create_time"`
}

type PrizeDao struct {
	gobase.TableHelper
}

func (dao *PrizeDao) GetSome(list *[]Prize) error {
	return dao.GetAll(list, "id, name, create_time", 5)
}