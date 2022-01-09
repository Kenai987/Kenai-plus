package main

import (
	"Kenai-plus/models"
	_ "Kenai-plus/models"
	_ "Kenai-plus/routers"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("数据库连接开始。。。")
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1)/kenai?charset=utf8")
	if err != nil {
		logs.Info("创建数据库连接出错", err.Error())
	}
	//创建表
	orm.RegisterModel(new(models.Blog), new(models.User))
	//生成表
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Info("创建数据表出错", err.Error())
	}
}
func main() {
	beego.Run()
}
