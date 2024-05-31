package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cairomedeiros/todo-list/schemas"
)

func SendSuccess(w http.ResponseWriter, op string, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": fmt.Sprintf("operation from handler: %s successful", op),
		"data":    data,
	}
	json.NewEncoder(w).Encode(response)
}

func SendError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Type:    http.StatusText(code),
		Message: msg,
	})
}

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// task response
type CreateTaskResponse struct {
	Message string               `json:"message"`
	Data    schemas.TaskResponse `json:"data"`
}

type UpdateTaskResponse struct {
	Message string               `json:"message"`
	Data    schemas.TaskResponse `json:"data"`
}

type DeleteTaskResponse struct {
	Message string               `json:"message"`
	Data    schemas.TaskResponse `json:"data"`
}

type GetTaskByIdResponse struct {
	Message string               `json:"message"`
	Data    schemas.TaskResponse `json:"data"`
}

type ListTasksResponse struct {
	Message string                 `json:"message"`
	Data    []schemas.TaskResponse `json:"data"`
}

// subtask response
type CreateSubTaskResponse struct {
	Message string                  `json:"message"`
	Data    schemas.SubTaskResponse `json:"data"`
}

type UpdateSubTaskResponse struct {
	Message string                  `json:"message"`
	Data    schemas.SubTaskResponse `json:"data"`
}

type DeleteSubTaskResponse struct {
	Message string                  `json:"message"`
	Data    schemas.SubTaskResponse `json:"data"`
}

type GetSubTaskByIdResponse struct {
	Message string                  `json:"message"`
	Data    schemas.SubTaskResponse `json:"data"`
}

type ListSubTasksResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.SubTaskResponse `json:"data"`
}
