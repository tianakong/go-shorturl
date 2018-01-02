package main

import (
	_ "shorturl/routers"
	"github.com/astaxie/beego"
	"shorturl/models"
)

func main() {
	models.Init()
	beego.Run()
}

