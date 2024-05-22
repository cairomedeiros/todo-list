package schemas

import (
	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	Title       string
	Description string
}
