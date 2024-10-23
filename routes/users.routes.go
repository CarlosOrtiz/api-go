package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/CarlosOrtiz/api-go/config"
	"github.com/CarlosOrtiz/api-go/config/dto"
	"github.com/CarlosOrtiz/api-go/config/utils"
	"github.com/CarlosOrtiz/api-go/database"
	"github.com/CarlosOrtiz/api-go/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	var totalItems int64

	/* pageStr := r.URL.Query().Get("page")
	quantityStr := r.URL.Query().Get("quantity") */
	order := r.URL.Query().Get("order")
	name := r.URL.Query().Get("name")

	if order != "" && strings.ToLower(order) != "asc" && strings.ToLower(order) != "desc" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "INVALID_VALUE_ORDER",
			Message: "Valor invalido para el parametro 'order'. Debe ser 'asc' o 'desc'.",
		})
		return
	}

	/* page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {ƒ
		page = 1
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		quantity = 10
	} */

	//init query builder
	query := database.DB.Model(&models.User{})

	if name != "" {
		name = strings.ToLower(name)
		query = query.Where("LOWER(name) ILIKE ?", "%"+name+"%")
	}

	query.Count(&totalItems)

	var err error
	query, page, quantity, err := utils.Paginate(w, r, query) // Modificamos aquí para recibir `quantity`
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "Error en los parámetros de paginación.",
			Message: err.Error(),
		})
		return
	}

	//offset := (page - 1) * quantity

	if order == "desc" {
		query = query.Order("created_at desc")
	} else {
		query = query.Order("created_at asc")
	}

	//query = query.Preload("Tasks").Limit(quantity).Offset(offset)

	// end query with query.Find get all users

	/* if err := query.Find(&users).Error; err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "Error SQL",
			Message: err.Error(),
		})
		return
	} */

	if err := query.Preload("Tasks").Find(&users).Error; err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "Error retrieving users",
			Message: err.Error(),
		})
		return
	}

	totalPages := int((totalItems + int64(quantity) - 1) / int64(quantity)) // Usamos `quantity`

	response := map[string]interface{}{
		"items": users,
		"meta": map[string]interface{}{
			"totalItems":   totalItems,
			"itemsPerPage": quantity,
			"totalPages":   totalPages,
			"currentPage":  page,
		},
	}

	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  response,
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

	database.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	//database.DB.Preload("Tasks").First(&user, params["id"])

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
