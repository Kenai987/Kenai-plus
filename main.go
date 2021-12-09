package main

import (
	_ "Kenai-plus/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

