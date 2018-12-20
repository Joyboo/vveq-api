package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		beego.Error("注册默认数据库失败: ", err)
	}
	orm.RegisterModel(new(User), new(Dau))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func Md5(str string) string {
	if str == "" {
		return ""
	}
	init := md5.New()
	init.Write([]byte(str))
	return fmt.Sprintf("%x", init.Sum(nil))
}
