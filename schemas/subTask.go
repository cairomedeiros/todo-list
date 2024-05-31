package schemas

import (
	"time"

	"gorm.io/gorm"
)

type SubTask struct {
	gorm.Model
	TaskID    uint
	Name      string
	Completed bool
}

type SubTaskResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	TaskID    uint      `json:"taskId"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
}
