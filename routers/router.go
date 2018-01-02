package routers

import (
	"shorturl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{}, "*:Index")
    beego.Router("/qrcode", &controllers.IndexController{}, "*:Qrcodeimg")
    beego.Router("/?:url", &controllers.IndexController{}, "*:Jump")

}
