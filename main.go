package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CarlosOrtiz/api-go/config/middleware"
	"github.com/CarlosOrtiz/api-go/database"
	"github.com/CarlosOrtiz/api-go/models"
	"github.com/CarlosOrtiz/api-go/routes"
	"github.com/gorilla/mux"
)

func main() {

	database.Connection()

	database.DB.AutoMigrate(models.Task{}, models.User{})

	router := mux.NewRouter()

	router.Use(middleware.ResponseJson)
	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/user", routes.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/{id}", routes.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/user/{id}", routes.DeleteUserHandler).Methods("DELETE")

	router.HandleFunc("/tasks", routes.GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", routes.CreateTask).Methods("GET")
	router.HandleFunc("/task", routes.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", routes.DeleteTask).Methods("DELETE")

	fmt.Println("Server started on port ", 7002)
	log.Fatal(http.ListenAndServe(":7002", router))
}
