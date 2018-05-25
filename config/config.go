package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/oujia/gobase"
)

const (
	LOG_KEY = "logstash:dw#1"
	LOG_SELF_CALL = "test_gobase_selfcall"
	LOG_MODULE_CALL = "test_gobase_modulecall"

	SERVER_ADDR = ":8089"

	GLOBAL_CONFIG_PATH = "/data/webapps/conf_v2/test/all.json"
)

var DB_LOCALHOST *sqlx.DB
var DB_HD *sqlx.DB
var GlobalConf *gobase.GlobalConf

func init()  {
	conf, err := gobase.LoadGlobalConf(GLOBAL_CONFIG_PATH)
	checkErr(err)
	GlobalConf = conf

	DB_LOCALHOST, err = gobase.NewDbClient("localhost", GlobalConf.DbInfo)
	checkErr(err)

	DB_HD, err = gobase.NewDbClient("dw_ka_hd", GlobalConf.DbInfo)
	checkErr(err)
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}