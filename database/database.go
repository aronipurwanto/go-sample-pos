package database

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:P@ssW0rd32!@tcp(localhost:3306)/sample_pos_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.Product{},
		&models.Customer{}) // Migrasi tabel product
	DB = database
}
