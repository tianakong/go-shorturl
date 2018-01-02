package controllers

import (
	"github.com/astaxie/beego"
	"shorturl/models"
	"github.com/astaxie/beego/orm"
	"github.com/skip2/go-qrcode"
	"github.com/astaxie/beego/cache"
)

var (
	urlcache cache.Cache
)

func init() {
	urlcache, _ = cache.NewCache("memory", `{"interval":0}`)
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {
	url := this.GetString("url")
	if url!="" {
		urlOne,err := models.UrlGetByUrl(url)
		if err == nil {
			urlcache.Put(urlOne.Short, url, 0)
			this.Data["shortUrl"] = beego.AppConfig.String("website") + urlOne.Short
			this.Data["longUrl"] = url
		} else {
			urlModel := new(models.Url)
			urlModel.Url = url
			id, _ := models.UrlAdd(urlModel)
			urlModel.Short = models.Generate(id)
			orm.NewOrm().Update(urlModel)
			urlcache.Put(urlModel.Short, url, 0)
			this.Data["shortUrl"] = beego.AppConfig.String("website") + urlModel.Short
			this.Data["longUrl"] = url
		}
	}
	this.TplName = "index/index.tpl";
}
func (this *IndexController) Qrcodeimg() {
	url := this.GetString("url")
	if url!="" {
		this.Ctx.ResponseWriter.Header().Set("Content-Type", "image/png")
		png, _ := qrcode.Encode(url, qrcode.Medium, 256)
		this.Ctx.ResponseWriter.Write(png)
		this.StopRun()
	}
}


func (this *IndexController) Jump() {
	url := this.Ctx.Input.Param(":url")
	longurl := urlcache.Get(url).(string)

	if longurl != "" {
		//model := new(models.Detail)
		//model.AccessIp = this.Ctx.Request.RemoteAddr
		this.Redirect(longurl, 302)
	}
	this.StopRun()
}
