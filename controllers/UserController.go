package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"zhenyan/lib/log"
	"zhenyan/models"
)

type UserController struct {
	web.Controller
}

var (
	logs = log.NewLogger()
)

func (u *UserController) HelloWorld() {
	users := models.GetUsersByName("s")
	valid := validation.Validation{}
	for _, u := range users {
		b, err := valid.Valid(u)
		if err != nil {
			logs.Error(err.Error())
		}
		if b {
			// validation does not pass
			// blabla...
			for _, err := range valid.Errors {
				logs.Debug(u.Name, u.ID, err.Key, err.Message)
			}
		}
	}
	fmt.Println(users)
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("models.GetUsersByName:json:error")
	}
	logs.Debug("HelloWorld")
	confStr, _ := web.AppConfig.String("test::k")
	logs.Debug("config:value:" + confStr)
	u.Ctx.WriteString("hello, world!!!!" + string(data) + confStr)
}
func (u *UserController) Hi() {
	appName, _ := web.AppConfig.String("appname")
	logs.Error("Fun is not impl" + appName)
	u.Ctx.WriteString("Fun is not impl")

}
