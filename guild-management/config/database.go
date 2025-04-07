package config

import (
	"guild-management/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Mimimutu88@tcp(127.0.0.1:3306)/guild_management?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Migrasi model Guild dan Member
	database.AutoMigrate(&models.Guild{}, &models.Member{})
	DB = database
	log.Println("Database connected successfully!")
}
