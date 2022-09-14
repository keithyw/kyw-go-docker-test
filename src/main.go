package main

import (
	"log"
	"net/http"
	"github.com/keithyw/kyw-go-docker-test/conf"
	"github.com/keithyw/kyw-go-docker-test/controllers"
	"github.com/keithyw/kyw-go-docker-test/database"
	"github.com/keithyw/kyw-go-docker-test/grpc"
	"github.com/keithyw/kyw-go-docker-test/repositories"
	"github.com/keithyw/kyw-go-docker-test/routers"
	"github.com/keithyw/kyw-go-docker-test/services"
)

var conn *database.MysqlDB
var repo repositories.UserRepository
var service services.UserService

func main() {
	config, _ := conf.NewConfig()
	conn = database.NewDatabase(config)
	defer conn.DB.Close()
	repo = repositories.NewUserRepository(conn)
	client := grpc.NewGrpcClient(config)
	defer client.Conn.Close()
	service = services.NewUserService(client, repo)
	controller := controllers.NewUserController(config, service)
	r := routers.NewUserRouter(controller)
	log.Fatal(http.ListenAndServe(config.Port, r.GetRouter()))
}