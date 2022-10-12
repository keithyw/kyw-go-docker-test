package routers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/keithyw/kyw-go-docker-test/controllers"
)

type ApiUserRouter struct {
	router *mux.Router
}

func NewApiUserRouter(controller controllers.ApiUserController) ApiUserRouter {
	router := mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/users", controller.Users).Methods(http.MethodGet)
	apiRouter.HandleFunc("/users/{id}", controller.Delete).Methods(http.MethodDelete)
	apiRouter.HandleFunc("/users/{id}", controller.Update).Methods(http.MethodPut)
	apiRouter.HandleFunc("/users", controller.Store).Methods(http.MethodPost)
	apiRouter.HandleFunc("/users/{id}", controller.View).Methods(http.MethodGet)
	// apiRouter.Use(mux.CORSMethodMiddleware(apiRouter))
	return ApiUserRouter{router}
}

func (r ApiUserRouter) GetRouter() *mux.Router {
	return r.router
}