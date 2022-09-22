package database

import (
	migration "day-7/database/migration"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
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

func InitDBMongo() *mongo.Client {
	confMongo := dbConfig{
		User: os.Getenv("DB_USERMONGO"),
		Pass: os.Getenv("DB_PASSMONGO"),
		Host: os.Getenv("DB_HOSTMONGO"),
		Port: os.Getenv("DB_PORTMONGO"),
		Name: os.Getenv("DB_NAMEMONGO"),
	}

	mongo := mongoConfig{dbConfig: confMongo}
	client := mongo.ConnectMongo()
	return client
}

func GetConnectionMongo() *mongo.Client {
	client := InitDBMongo()

	return client
}
