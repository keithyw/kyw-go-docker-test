package controllers

import (
	"fmt"
	"log"
	"strconv"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/keithyw/kyw-go-docker-test/conf"
	"github.com/keithyw/kyw-go-docker-test/models"
	"github.com/keithyw/kyw-go-docker-test/services"
)

type PageData struct {
	Title string
	Data interface{}
}

type UserController struct {
	config *conf.Config
	svc services.UserService
	templates *template.Template
}

func NewUserController(config *conf.Config, service services.UserService) UserController {
	templates := template.Must(template.ParseGlob("./views/*.html"))
	return UserController{config, service, templates}
}

func (uc *UserController) Home(w http.ResponseWriter, r *http.Request) {
	uc.render(w, "home.html", nil)
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var data PageData
	data.Title = "Create New User"
	data.Data = nil
	uc.render(w, "create.html", data)
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	err = uc.svc.DeleteUser(id)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/users", http.StatusFound)
}

func (uc *UserController) Store(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var u models.User
	u.Username = r.FormValue("username")
	u.Passwd = r.FormValue("passwd")
	u.Email = r.FormValue("email")
	u.FirstName = r.FormValue("first_name")
	u.LastName = r.FormValue("last_name")
	existingUser, _ := uc.svc.FindUserByName(u.Username)
	if existingUser != nil {
		http.Error(w, "User already exist by that name", 500)
		return
	}
	newUser, err := uc.svc.CreateUser(u)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r,fmt.Sprintf("/users/%d", newUser.ID), http.StatusFound)
}

func (uc *UserController) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var data PageData
	data.Title = "Edit User"
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	user, err := uc.svc.FindUserById(id)
	if err != nil {
		panic(err)
	} else {
		data.Data = user
	}
	uc.render(w, "edit.html", data)
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		panic(err)
	}
	var u models.User
	u.Username = r.FormValue("username")
	u.Email = r.FormValue("email")
	u.FirstName = r.FormValue("first_name")
	u.LastName = r.FormValue("last_name")
	log.Println("username " + u.Username)
	log.Printf("id %d", id)
	existingUser, _ := uc.svc.FindUserById(id)

	if existingUser == nil {
		http.Error(w, "User does not exist", 500)
		return
	}
	_, err = uc.svc.UpdateUser(id, u)
	if err != nil {
		panic(err)
	}
	log.Println("updated user")
	http.Redirect(w, r,fmt.Sprintf("/users/%d", id), http.StatusFound)
}

func (uc *UserController) Users(w http.ResponseWriter, r *http.Request) {
	var data PageData
	data.Title = "Users"
	users, err := uc.svc.GetAllUsers()
	if err != nil {
		log.Printf("No data but have error %s", err)
		data.Data = nil
	} else {
		data.Data = users
	}
	uc.render(w, "users.html", data)
}

func (uc *UserController) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var data PageData
	data.Title = "User"
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	user, err := uc.svc.FindUserById(id)
	if err != nil {
		panic(err)
	}
	data.Data = user
	uc.render(w, "view.html", data)
}

func (uc *UserController) render(w http.ResponseWriter, templateName string, data interface{}) {
	if err := uc.templates.ExecuteTemplate(w, templateName, data); err != nil {
		panic(err)
	}
}