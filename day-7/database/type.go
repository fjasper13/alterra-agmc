package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	dbConfig struct {
		Host string
		User string
		Pass string
		Port string
		Name string
	}

	mysqlConfig struct {
		dbConfig
	}

	mongoConfig struct {
		dbConfig
	}
)

func (conf mysqlConfig) Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Name,
	)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func (conf mongoConfig) ConnectMongo() *mongo.Client {

	url := fmt.Sprintf("mongodb://%s:%s",
		conf.Host,
		conf.Port,
	)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}

	return client

}
