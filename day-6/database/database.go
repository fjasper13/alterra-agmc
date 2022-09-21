package database

import (
	migration "day-6/database/migration"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	conf := dbConfig{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}

	mysql := mysqlConfig{dbConfig: conf}

	mysql.Connect()

	migration.InitMigrate(DB)
}

func GetConnection() *gorm.DB {
	if DB == nil {
		InitDB()
	}
	return DB
}