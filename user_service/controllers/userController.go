package controllers

import (
	"database/sql"
	"encoding/json"
	"microservice_golang/user_service/models"
	"microservice_golang/user_service/repositories"
	"microservice_golang/user_service/services"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	router.HandleFunc("/users", createUser(userService)).Methods("POST")
	router.HandleFunc("/users", getAllUser(userService)).Methods("GET")
}

func createUser(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		err := userService.CreateUser(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

func getAllUser(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.GetAllUser()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}
