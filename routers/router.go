package routers

import (
	"golang_basic_rest_server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
    beego.Router("/test", &controllers.TestController{})
	
}
