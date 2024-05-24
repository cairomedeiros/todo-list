package handler

import "time"

type CreateTaskRequest struct {
	Id          uint       `json:"id" gorm:"primaryKey"`
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
	Id        uint   `json:"id" gorm:"primaryKey"`
	TaskID    uint   `json:"taskId"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type UpdateSubTaskRequest struct {
	TaskID    uint   `json:"taskId"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
