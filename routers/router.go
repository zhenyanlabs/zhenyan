package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"zhenyan/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.UserController{})
}
