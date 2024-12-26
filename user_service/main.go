package main

import (
	"log"
	"microservice_golang/user_service/config"
	"microservice_golang/user_service/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := mux.NewRouter()
	controllers.RegisterUserRoutes(router, db)

	log.Println("User service running on part 8000")
	http.ListenAndServe(":8000", router)
}
