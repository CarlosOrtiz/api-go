package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/CarlosOrtiz/api-go/config"
	"github.com/CarlosOrtiz/api-go/config/dto"
	"github.com/CarlosOrtiz/api-go/database"
	"github.com/CarlosOrtiz/api-go/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)

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
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "USER_NOT_FOUND",
		})
		return
	}

	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  &user,
	})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	user.Name = strings.ToUpper(user.Name)
	user.Lastname = strings.ToUpper(user.Lastname)
	user.Email = strings.ToLower(user.Email)

	newUser := database.DB.Create(&user)
	err := newUser.Error
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  &user,
	})
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "INVALID_PAYLOAD",
		})
		return
	}

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "INVALID_USER_ID",
		})
		return
	}

	var user models.User
	database.DB.First(&user, userID)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "USER_NOT_FOUND",
		})
		return
	}

	switch {
	case userDTO.Name != "":
		user.Name = strings.ToUpper(userDTO.Name)
	case userDTO.LastName != "":
		user.Lastname = strings.ToUpper(userDTO.LastName)
	case userDTO.Email != "":
		user.Email = strings.ToLower(userDTO.Email)
	}

	database.DB.Save(&user)

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
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "USER_NOT_FOUND",
		})
		return
	}

	//database.DB.Delete(&user)
	database.DB.Unscoped().Delete(&user)

	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  &user,
		Message: "Usuario eliminado exitosamente",
	})
}
