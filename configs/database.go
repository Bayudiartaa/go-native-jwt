package configs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  "log"
  "go-native-jwt/models"
)

var DB *gorm.DB 

func ConnectDB() {
  db, err := gorm.Open(postgres.Open("postgresql://postgres:password@localhost:5432/go_native?sslmode=disable&TimeZone=Asia/Jakarta"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&models.User{}, &models.Product{})

  DB = db
  log.Println("Databases connected")
}