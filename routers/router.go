package routers

import (
	"pachong/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/reptile", &controllers.CrawlMovieController{}, "*:CrawlMovie")
}
