package main

import (
	"log"
	"net/http"
	//github.com/astaxie/beego
	"github.com/beego/beego/v2/server/web"
	"github.com/keithyw/kyw-go-docker-test/conf"
	"github.com/keithyw/kyw-go-docker-test/controllers"
	"github.com/keithyw/kyw-go-docker-test/database"
	"github.com/keithyw/kyw-go-docker-test/repositories"
	"github.com/keithyw/kyw-go-docker-test/routers"
	"github.com/keithyw/kyw-go-docker-test/services"
)

var conn *database.MysqlDB
var repo repositories.UserRepository
var service services.UserService

type User struct {
	ID int64
	Username string
}

type MainController struct {
	web.Controller
}

func main() {
	config, _ := conf.NewConfig()
	conn = database.NewDatabase(config)
	defer conn.DB.Close()
	repo = repositories.NewUserRepository(conn)
	service = services.NewUserService(repo)
	controller := controllers.NewUserController(config, service)
	r := routers.NewUserRouter(controller)
	log.Fatal(http.ListenAndServe(config.Port, r.GetRouter()))
	// web.Router("/home", &MainController{})
	// web.Run()
}

func (this *MainController) Get() {
	
	users, err := service.GetAllUsers()
	if err != nil {
		panic(err)
	}
	
	this.Data["users"] = users
	this.Data["name"] = "keith"
	this.TplName = "results.html"
}