package routes

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBinstance() *gorm.DB {
	dsn := "root:kl18jda183079@tcp(localhost:3306)/calorie_tracker?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Failed to conect to the database: ",err)
	}

	sqlDB, err := db.DB()
	if err != nil{
		log.Fatal("Failed to get DB instance: ",err)
	}
	err = sqlDB.Ping()
	if err != nil{
		log.Fatal("Failed to ping the database: ",err)
	}

	fmt.Println("Connected to MySQL database")
	return db
}

func OpenCollection(db *gorm.DB,tableName string)*gorm.DB{
	return db.Table(tableName)
}
