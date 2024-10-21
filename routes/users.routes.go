package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarlosOrtiz/api-go/config"
	"github.com/CarlosOrtiz/api-go/database"
	"github.com/CarlosOrtiz/api-go/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  &users,
	})
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	database.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "USER_NOT_FOUND",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  &user,
	})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	newUser := database.DB.Create(&user)
	err := newUser.Error
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  &user,
	})
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(&user)

	database.DB.First(&user)

	if user.ID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "USER_NOT_FOUND",
		})
		return
	}

	database.DB.Model(&user).Updates(&user)

	database.DB.First(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  &user,
	})
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	database.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "USER_NOT_FOUND",
		})
		return
	}

	//database.DB.Delete(&user)
	database.DB.Unscoped().Delete(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  &user,
		Message: "Usuario eliminado exitosamente",
	})
}
