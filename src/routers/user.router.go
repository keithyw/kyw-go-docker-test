package routers

import (
	// "log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/keithyw/kyw-go-docker-test/controllers"
)

type UserRouter struct {
	router *mux.Router
}

// func setHeader(header, value string, handle http.Handler) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		w.Header().Set(header, value)
// 		log.Println("setting header " + header)
// 		log.Println("header is " + value)
// 		handle.ServeHTTP(w, req)
// 	}
// }

func NewUserRouter(controller controllers.UserController) UserRouter {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/home", controller.Home).Methods(http.MethodGet)
	router.HandleFunc("/users", controller.Users).Methods(http.MethodGet)
	router.HandleFunc("/create", controller.Create).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}/delete", controller.Delete).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}/edit", controller.Edit).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controller.Update).Methods(http.MethodPost)
	router.HandleFunc("/user", controller.Store).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", controller.View).Methods(http.MethodGet)
	
	return UserRouter{router}
}

func (r UserRouter) GetRouter() *mux.Router {
	return r.router
}