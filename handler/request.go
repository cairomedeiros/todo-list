package handler

import (
	"fmt"
	"time"
)

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateTaskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
	Completed   bool       `json:"completed"`
}

type UpdateTaskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
	Completed   bool       `json:"completed"`
}

type CreateSubTaskRequest struct {
	TaskID    uint   `json:"taskId"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type UpdateSubTaskRequest struct {
	TaskID    uint   `json:"taskId"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

//user

type RegisterRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type LoginRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
