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
