package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/keithyw/kyw-go-docker-test/conf"
	"github.com/keithyw/kyw-go-docker-test/models"
	"github.com/keithyw/kyw-go-docker-test/services"
)

type ApiUserController struct {
	config *conf.Config
	svc services.UserService
}

type ApiUserResponse struct {
	IsSuccess bool `json:"is_success,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewApiUserController(config *conf.Config, service services.UserService) ApiUserController {
	return ApiUserController{config, service}
}

func (uc *ApiUserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	err = uc.svc.DeleteUser(id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ApiUserResponse{IsSuccess: true})
}

func (uc *ApiUserController) Store(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	var u models.User
	err = json.Unmarshal(b, &u)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(ApiUserResponse{IsSuccess: false, Message: err.Error()})
		return
	}
	
	existingUser, _ := uc.svc.FindUserByName(u.Username)
	if existingUser != nil {
		json.NewEncoder(w).Encode(ApiUserResponse{IsSuccess: false, Message: "User already exist by that name"})
		return
	}
	newUser, err := uc.svc.CreateUser(u)
	if err != nil {
		json.NewEncoder(w).Encode(ApiUserResponse{IsSuccess: false, Message: "Failed creating user: " + err.Error()})
		return
	}
	jsonString, err := json.Marshal(newUser)
	if err != nil {
		json.NewEncoder(w).Encode(ApiUserResponse{IsSuccess: false, Message: "Failed marshaling data: " + err.Error()})
		return
	}
	w.Write(jsonString)
}

func (uc *ApiUserController) Update(w http.ResponseWriter, r *http.Request) {
	var u models.User
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		panic(err)
	}
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &u)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	
	existingUser, _ := uc.svc.FindUserById(id)

	if existingUser == nil {
		http.Error(w, "User does not exist", 500)
		return
	}
	updatedUser, err := uc.svc.UpdateUser(id, u)
	if err != nil {
		panic(err)
	}
	jsonString, err := json.Marshal(updatedUser)
	if err != nil {
		panic(err)
	}
	log.Println("updated user")
	w.Write(jsonString)
}

func (uc *ApiUserController) Users(w http.ResponseWriter, r *http.Request) {
	users, err := uc.svc.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Printf("No data but have error %s", err)
		json.NewEncoder(w).Encode(ApiUserResponse{IsSuccess: false, Message: fmt.Sprintf("No data %s", err)})
		return
	} 
	jsonString, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.Write(jsonString)
}

func (uc *ApiUserController) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	user, err := uc.svc.FindUserById(id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonString, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Write(jsonString)
}