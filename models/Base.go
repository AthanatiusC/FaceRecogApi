package models

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Payload struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var db *mongo.Database //database
// var db_sas *gorm.DB    //database
// var db_master *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	db = client.Database("smart_school")
	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASS")
	// db_name := os.Getenv("DB_NAME")
	// db_name_sas := os.Getenv("DB_SEMAYA_SAS")
	// db_name_master := os.Getenv("DB_SEMAYA_MASTER")
	// dbHost := os.Getenv("db_host")
	// dbURI := fmt.Sprintf("host=localhost port=5433 user=%s dbname=%s password=%s sslmode=disable", username, db_name, password)
	// dbURI := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, db_name)
	// dbURISas := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, db_name_sas)
	// dbURIMaster := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, db_name_master)

	// conn, err := gorm.Open("postgres", dbURI)
	// conn, err := gorm.Open("mysql", dbURI)

	// connSas, err := gorm.Open("mysql", dbURISas)
	// connMaster, err := gorm.Open("mysql", dbURIMaster)

	// if err != nil {
	// 	fmt.Print(err)
	// }

	// db = conn.LogMode(true)
	// db.Set("gorm:auto_preload", true)
	// // db_sas = connSas
	// db_master = connMaster

	// db.Debug().AutoMigrate(&User{}, &Camera{}, &Attendance{}, &Role{}) //Database migration
}

// GetDB : returns a handle to the DB object
func GetDB(whatDB string) *mongo.Database {
	if whatDB == "main" {
		return db
	} else {
		return db
	}
	return db
}
