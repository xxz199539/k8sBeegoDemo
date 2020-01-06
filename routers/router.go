package routers

import (
	"github.com/astaxie/beego"
	"k8sBeegoDemo/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Include(&controllers.UserController{})
}
