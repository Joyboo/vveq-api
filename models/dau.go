package models

import (
	"bytes"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

type Dau struct {
	Id      int64
	Uid     int64
	Ip      int64
	Instime int64
}

var tablenameDau = "dau"

func newDau() *Dau {
	return &Dau{}
}

func (d *Dau) AddDau() (int64, error) {
	d.Instime = time.Now().Unix()
	// TODO 走redis队列
	return orm.NewOrm().Insert(d)
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
