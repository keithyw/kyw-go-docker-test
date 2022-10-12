package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	// controller := controllers.NewUserController(config, service)
	apiController := controllers.NewApiUserController(config, service)
	r := routers.NewApiUserRouter(apiController)
	credentials := handlers.AllowCredentials()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "DELETE", "PUT", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	// r := routers.NewUserRouter(controller)
	log.Fatal(http.ListenAndServe(config.Port, handlers.CORS(credentials, headersOk, methods, origins)(r.GetRouter())))
}