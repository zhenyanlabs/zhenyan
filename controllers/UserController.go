package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"zhenyan/lib/log"
	"zhenyan/models"
)

type UserController struct {
	web.Controller
}

func (u *UserController) HelloWorld() {
	users := models.GetUsersByName("s")
	fmt.Println(users)
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("models.GetUsersByName:json:error")
	}
	log.NewLogger().Debug("HelloWorld")
	u.Ctx.WriteString("hello, world!!!!" + string(data))
}
