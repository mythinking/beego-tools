package routers

import (
	"beego-tools/controllers"
	"beego-tools/controllers/crawlers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/crawlers",
			beego.NSNamespace("/eshop",
				beego.NSRouter("/", &crawlers.TmallController{}, "get:Index;post:CreateJob"),
			),
			beego.NSRouter("/log", &controllers.CrawlerController{}, "get:Log"),
		),
	)
	beego.AddNamespace(ns)
}
