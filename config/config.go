package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = InitializePostgreSQL()

	if err != nil {
		return nil
	}

	return nil
}

func GetPostgreSQL() *gorm.DB {
	return db
}
