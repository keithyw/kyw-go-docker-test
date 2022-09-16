package routers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/keithyw/kyw-go-docker-test/controllers"
)

type UserRouter struct {
	router *mux.Router
}

func NewUserRouter(controller controllers.UserController) UserRouter {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/home", controller.Home).Methods(http.MethodGet)
	router.HandleFunc("/users", controller.Users).Methods(http.MethodGet)
	router.HandleFunc("/users/create", controller.Create).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}/delete", controller.Delete).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}/edit", controller.Edit).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controller.Update).Methods(http.MethodPost)
	router.HandleFunc("/users", controller.Store).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", controller.View).Methods(http.MethodGet)
	
	return UserRouter{router}
}

func (r UserRouter) GetRouter() *mux.Router {
	return r.router
}