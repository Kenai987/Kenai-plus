package controllers

import (
	"Kenai-plus/models"
	"github.com/beego/beego/v2/adapter/logs"
	orm2 "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"

}

func (c *LoginController) Post() {
	name := c.GetString("name")
	pwd := c.GetString("pwd")

	if name == "" || pwd == "" {
		c.Redirect("/", 302)
	}

	orm := orm2.NewOrm()
	user := models.User{}
	user.Name = name
	user.Pwd = pwd
	err := orm.Read(&user)
	if err != nil {
		logs.Info(err.Error())
	}
	//设置session
	err = c.SetSession(user.Name, 1)
	if err != nil {
		logs.Info("session设置失败")
	}
	c.Data["data"] = user.Name
	c.TplName = "internal.html"

}
