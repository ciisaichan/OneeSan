package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type IllustInfo struct {
	Id          int
	Pid         int
	Title       string
	Author      string
	OriginalUrl string
	MasterUrl   string
}

func init() {
	db_addr := beego.AppConfig.String("db_addr")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	orm.RegisterDataBase("default", "mysql", db_user+":"+db_pass+"@tcp("+db_addr+")/"+db_name+"?charset=utf8mb4")
	orm.RegisterModel(new(IllustInfo))
	orm.RunSyncdb("default", false, true)
}
