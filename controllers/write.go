package controllers

import (
	"Kenai-plus/models"
	"github.com/beego/beego/v2/adapter/logs"
	orm2 "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type WriterController struct {
	beego.Controller
}

func (c *WriterController) Get() {
	c.TplName = "write.html"

}

func (c *WriterController) Post() {
	title := c.GetString("title")
	content := c.GetString("content")

	if title == "" || content == "" {
		c.Redirect("/", 302)
	}

	orm := orm2.NewOrm()
	blog := models.Blog{}
	blog.Title = title
	blog.Content = content

	_, err := orm.Insert(&blog)
	if err != nil {
		logs.Info("文章创建失败")
	}
	c.Data["data"] = blog.Title
	c.TplName = "hello.html"

}
