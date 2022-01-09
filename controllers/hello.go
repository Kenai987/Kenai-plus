package controllers

import (
	"Kenai-plus/models"
	"github.com/beego/beego/v2/adapter/logs"
	orm2 "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	// 1、ORM对象
	orm := orm2.NewOrm()
	// 2、结构体对象
	user := models.User{}
	// 3、结构体对象赋值
	user.Name = "libai"
	// 4、插入
	insert, err := orm.Insert(&user)
	if err != nil {
		logs.Info(err.Error())
	}
	logs.Info(insert)

	// 查询
	// 1、ORM对象
	// 2、查询的对象
	// 3、构建条件
	// 4、查询
	err = orm.Read(&user)
	if err != nil {
		logs.Info("查询失败", err)
	}
	logs.Info("查询成功", user)

	c.Data["data"] = insert
	c.TplName = "hello.html"
}

func (c *HelloController) Post() {
	c.Data["data"] = "你点击了一次"
	c.TplName = "hello.html"
}
