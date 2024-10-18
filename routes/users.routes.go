package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CarlosOrtiz/api-go/config"
	"github.com/CarlosOrtiz/api-go/database"
	"github.com/CarlosOrtiz/api-go/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	newUser := database.DB.Create(&user)
	err := newUser.Error
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
	}

	response := config.BasicResponse{
		Success: true,
		Detail:  &user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update "))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete "))
}
