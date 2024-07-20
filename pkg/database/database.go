package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kanban-board-app/internal/models"
	"log"
)

var DB *gorm.DB

func Connect(username, password, dbName, host string, port int) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
}

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Board{}, &models.List{}, &models.Card{})
}
