package handler

import (
	"github.com/cairomedeiros/todo-list/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	db = config.GetPostgreSQL()
}
