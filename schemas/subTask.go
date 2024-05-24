package schemas

import (
	"gorm.io/gorm"
)

type SubTask struct {
	gorm.Model
	TaskID    uint
	Name      string
	Completed bool
}
