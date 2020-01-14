package db

import (
	"conscience-backend/config"
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var DB *gorm.DB
var Mongo *mongo.Client

func Init(conf config.DB) (err error) {
	Mongo, err = initMongo(conf.Mongo.Uri)
	DB, err = gorm.Open("mysql", conf.Mysql.DSN)
	return
}

func Close() {
	DB.Close()
}

func initMongo(uri string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
