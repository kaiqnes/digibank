package database

import (
	"digibank/internal/domain/entities"
	"digibank/internal/frameworks/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewSession() *gorm.DB {
	dns := getDNS()
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database - err %v", err)
	}
	migrateModels(db)
	return db
}

func migrateModels(db *gorm.DB) {
	db.AutoMigrate(&entities.Account{})
	db.AutoMigrate(&entities.OperationType{})
	db.AutoMigrate(&entities.Transaction{})
}

func getDNS() string {
	user := utils.GetEnvOrDefault("DB_USER", "")
	pass := utils.GetEnvOrDefault("DB_PASS", "")
	host := utils.GetEnvOrDefault("DB_HOST", "")
	port := utils.GetEnvOrDefault("DB_PORT", "")
	name := utils.GetEnvOrDefault("DB_NAME", "")

	dnsTemplate := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf(dnsTemplate, user, pass, host, port, name)
}
