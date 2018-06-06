package model

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
	Id int `db:"id,ai" redis:"id"`
	Hash string `db:"hash" redis:"hash"`
	Title string `db:"title" redis:"title"`
}

type EpisodeDao struct {
	gobase.TableHelper
}

