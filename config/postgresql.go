package config

import (
	"fmt"

	"github.com/cairomedeiros/todo-list/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=tododb port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Task{})
	if err != nil {
		fmt.Println("Error during migration:", err)
		return nil, err
	}

	return db, nil
}
