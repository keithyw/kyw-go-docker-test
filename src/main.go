package main

import (
	//github.com/astaxie/beego
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Router("/home", &MainController{})
	web.Run()
}

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	this.Data["name"] = "keith"
	this.TplName = "results.html"
}