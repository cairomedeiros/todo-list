package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	DueDate     *time.Time
	Completed   bool
	SubTasks    *[]SubTask `gorm:"foreignKey:TaskID"`
}

type TaskResponse struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   time.Time  `json:"deteledAt,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
	Completed   bool       `json:"completed"`
}
