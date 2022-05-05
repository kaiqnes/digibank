package database

import (
	"database/sql"
	"digibank/internal/domain/entities"
	"digibank/internal/frameworks/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewSession() *gorm.DB {
	createDatabase()

	dns := getFullDNS()
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database - err %v", err)
	}
	migrateModels(db)
	return db
}

func createDatabase() {
	db, err := sql.Open("mysql", getShortDNS())
	if err != nil {
		log.Fatalf("failed to connect database - err %v", err)
	}
	defer db.Close()

	name := utils.GetEnvOrDefault("DB_NAME", "")
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", name))
	if err != nil {
		log.Fatalf("failed to create %s database - err %v", name, err)
	}
}

func migrateModels(db *gorm.DB) {
	db.AutoMigrate(&entities.Account{})
	db.AutoMigrate(&entities.OperationType{})
	db.AutoMigrate(&entities.Transaction{})
}

func getFullDNS() string {
	user := utils.GetEnvOrDefault("DB_USER", "")
	pass := utils.GetEnvOrDefault("DB_PASS", "")
	host := utils.GetEnvOrDefault("DB_HOST", "")
	port := utils.GetEnvOrDefault("DB_PORT", "")
	name := utils.GetEnvOrDefault("DB_NAME", "")

	dnsTemplate := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf(dnsTemplate, user, pass, host, port, name)
}

func getShortDNS() string {
	user := utils.GetEnvOrDefault("DB_USER", "")
	pass := utils.GetEnvOrDefault("DB_PASS", "")
	host := utils.GetEnvOrDefault("DB_HOST", "")
	port := utils.GetEnvOrDefault("DB_PORT", "")

	dnsTemplate := "%s:%s@tcp(%s:%s)/"
	return fmt.Sprintf(dnsTemplate, user, pass, host, port)
}
