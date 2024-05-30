package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cairomedeiros/todo-list/schemas"
)

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
