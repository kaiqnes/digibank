package database

import (
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
	return db
}

func getDNS() string {
	// TODO: Remove default values
	user := utils.GetEnvOrDefault("DB_USER", "root")
	pass := utils.GetEnvOrDefault("DB_PASS", "my-secret-pw")
	host := utils.GetEnvOrDefault("DB_HOST", "127.0.0.1")
	port := utils.GetEnvOrDefault("DB_PORT", "3306")
	name := utils.GetEnvOrDefault("DB_NAME", "challenge")

	dnsTemplate := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf(dnsTemplate, user, pass, host, port, name)
}
