package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CarlosOrtiz/api-go/config"
	"github.com/CarlosOrtiz/api-go/database"
	"github.com/CarlosOrtiz/api-go/models"
	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	database.DB.Find(&tasks)

	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  tasks,
	})
}

func GetOneTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	database.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "TASK_NOT_FOUND",
		})
		return
	}

	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  &task,
	})
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	newTask := database.DB.Create(&task)
	err := newTask.Error
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "TASK_NOT_CREATED",
		})
		return
	}
	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  task,
	})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	database.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(config.BasicResponse{
			Success: false,
			Detail:  "TASK_NOT_FOUND",
		})
		return
	}

	database.DB.Unscoped().Delete(&task)

	json.NewEncoder(w).Encode(config.BasicResponse{
		Success: true,
		Detail:  &task,
		Message: "Tarea eliminada exitosamente",
	})
}
