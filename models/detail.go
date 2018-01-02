package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

/**
 * 跳转记录表
 */
type Detail struct {
	Id uint64
	UrlId int
	AccessTime time.Time
	Device string `orm:"size(20)"`
	Browser string `orm:"size(20)"`
	ResolutionRatio string `orm:"size(20)"`  //分辨率
	AccessIp string `orm:"size(20)"`
	Source string
}

func (this *Detail) TableName() string {
	return TableName("detail")
}

func (this *Detail) DetailAdd(detail *Detail) (int64, error) {
	detail.AccessTime = time.Now()
	return orm.NewOrm().Insert(detail)
}
