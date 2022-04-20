package main

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "zhenyan/routers"
)

//http://127.0.0.1:8081/user/helloworld

func main() {

	beego.BConfig.Listen.EnableAdmin = true
	beego.BConfig.Listen.AdminAddr = "localhost"
	beego.BConfig.Listen.AdminPort = 8088
	orm.Debug = true
	beego.Run()
}
