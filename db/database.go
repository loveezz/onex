package db

import (
	"log"
	"one/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn := "host=localhost user=postgres password=123 dbname=qwe port=5432 sslmode=disable TimeZone=UTC"
	
	log.Printf("Попытка подключения к PostgreSQL с DSN: %s", dsn)

	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[-] Ошибка подключения к базе данных: %v", err)
	}


	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("[-] Ошибка получения DB интерфейса: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)           
	sqlDB.SetMaxOpenConns(100)          
	sqlDB.SetConnMaxLifetime(time.Hour)  

	
	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("[-] Ошибка автомиграции: %v", err)
	}

	log.Println("[+] Успешное подключение к PostgreSQL и миграция моделей")
}