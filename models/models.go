package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func Init() {
	host := beego.AppConfig.String("db.host")
	user := beego.AppConfig.String("db.user")
	pass := beego.AppConfig.String("db.password")
	name := beego.AppConfig.String("db.name")
	port := beego.AppConfig.String("db.port")
	charset := beego.AppConfig.String("db.charset")

	dns := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=" + charset

	orm.RegisterDataBase("default", "mysql", dns, 30)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	orm.RegisterModel(new(Url), new(Detail))
	orm.RunSyncdb("default", false, true)
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
