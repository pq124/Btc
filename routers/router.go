package routers

import (
	"BtWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.RegisterConntroller{})
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/bcsearch",&controllers.SearchController{})
}
