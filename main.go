package main

import (
	"net/http"

	"github.com/CarlosOrtiz/api-go/database"
	"github.com/CarlosOrtiz/api-go/models"
	"github.com/CarlosOrtiz/api-go/routes"
	"github.com/gorilla/mux"
)

func main() {

	database.Connection()

	database.DB.AutoMigrate(models.Task{})
	database.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/user", routes.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user", routes.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/user/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":7002", router)
}
