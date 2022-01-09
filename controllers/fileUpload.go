package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type FileUploadController struct {
	beego.Controller
}

func (c *FileUploadController) Get() {
	c.TplName = "fileUpload.html"
}

func (c *FileUploadController) Post() {
	file, header, err := c.GetFile("file")

	defer file.Close()
	if err != nil {

	} else {
		c.SaveToFile("file", "./static/img/"+header.Filename)
	}
}
