package schemas

import (
	"time"

	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	Id           string
	Title        string
	Description  string
	CreationDate time.Time
}
