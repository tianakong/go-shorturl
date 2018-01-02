package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

/**
 * 网址对照表
 */
type Url struct {
	Id         int
	Url        string `orm:"unique;size(1000)"`
	Short      string `orm:"unique;size(6)"`
	CreateTime time.Time
	JumpNumber int
}

func (this *Url) TableName() string {
	return TableName("url")
}

func UrlAdd(url *Url) (int64, error) {
	url.CreateTime = time.Now()
	url.JumpNumber = 0
	return orm.NewOrm().Insert(url)
}

func UrlGetByUrl(url string) (*Url, error) {
	urlModel := new(Url)
	err := orm.NewOrm().QueryTable(urlModel).Filter("Url", url).One(urlModel)
	if err == nil {
		return urlModel,nil
	}
	return nil,err
}

func Generate(id int64) (tiny string) {
	num := id
	alpha := merge(getRange(48, 57), getRange(65, 90))
	alpha = merge(alpha, getRange(97, 122))
	if num < 62 {
		tiny = string(alpha[num])
		return tiny
	} else {
		var runes []rune
		runes = append(runes, alpha[num%62])
		num = num / 62
		for num >= 1 {
			if num < 62 {
				runes = append(runes, alpha[num-1])
			} else {
				runes = append(runes, alpha[num%62])
			}
			num = num / 62

		}
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		tiny = string(runes)
		return tiny
	}
	return tiny
}

func getRange(start, end rune) (ran []rune) {
	for i := start; i <= end; i++ {
		ran = append(ran, i)
	}
	return ran
}

func merge(a, b []rune) []rune {
	c := make([]rune, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return c
}
