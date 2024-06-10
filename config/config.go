package config

import (
	"log"
	"os"

	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	SecretKey []byte
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

func LoadConfig() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}
	SecretKey = []byte(secret)
}
