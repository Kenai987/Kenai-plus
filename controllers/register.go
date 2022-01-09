package controllers

import (
	"Kenai-plus/models"
	"github.com/beego/beego/v2/adapter/logs"
	orm2 "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"

}

func (c *RegisterController) Post() {
	name := c.GetString("name")
	pwd := c.GetString("pwd")

	if name == "" || pwd == "" {
		c.Redirect("/", 302)
	}

	orm := orm2.NewOrm()
	user := models.User{}
	user.Name = name
	user.Pwd = pwd
	insert, err := orm.Insert(&user)
	if err != nil {
		logs.Info(err.Error())
	}
	c.Data["data"] = insert
	c.TplName = "hello.html"

}
