package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"zhenyan/lib/log"
)

var (
	logs = log.NewLogger()
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/zhenyan?charset=utf8")

	// need to register models in init
	orm.RegisterModel(new(User))
}

// User -
type User struct {
	ID   int    `orm:"column(id)"`
	Name string `orm:"column(name)" valid:"Required;Match(/^sam.*/)"`
}

func (u *User) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "sam") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 sam")
	}
}

func GetUsersByName(name string) []User {
	var users []User
	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("user.name",
		"user.name").
		From("user").
		Where("name like ?").
		OrderBy("name").Desc().
		Limit(10).Offset(0)
	// 导出 SQL 语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	sqlLikeName := "%" + name + "%"
	o.Raw(sql, sqlLikeName).QueryRows(&users)
	logs.Debug("model:" + sqlLikeName)
	return users
}
