package models

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Pwd  string
	Age  string
}

func init() {
	//设置数据库基本信息
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/kenai?charset=utf8")
	if err != nil {
		logs.Info("数据库链接出错", err.Error())
	}
	//创建表
	orm.RegisterModel(new(User))
	//生成表
	err = orm.RunSyncdb("default", true, true)
	if err != nil {
		logs.Info("创建数据库出错", err.Error())
	}
}
