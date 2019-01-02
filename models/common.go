package models

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Model interface {
	Query() orm.QuerySeter
	Add() (int64, error)
}

const TimeFormart = "2006-01-02 15:04:05"

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
	orm.RegisterModel(
		NewUser(),
		NewDau(),
		NewCate(),
		NewTheme(),
		NewTag(),
		NewComment(),
		NewLike(),
	)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	// 分区数据表，不支持自动建表
	/*err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
	orm.RunCommand()*/
}

func Md5(str string) string {
	if str == "" {
		return ""
	}
	init := md5.New()
	init.Write([]byte(str))
	return fmt.Sprintf("%x", init.Sum(nil))
}

// 字符串转换整型
func IpString2Int(ipstring string) int64 {
	ipSegs := strings.Split(ipstring, ".")
	ipInt := 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return int64(ipInt)
}

// 整型转换成字符串
func IpInt2String(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}
