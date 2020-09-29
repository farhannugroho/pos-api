package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pos_api/model"
)

var DB *gorm.DB

func InitDb() {
	// Configure DB Connection
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		Config.DbHost, Config.DbUser, Config.DbPassword, Config.DbName, Config.DbPort, Config.DbSslMode, Config.DbTimeZone)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("config.InitDb, Failed to connect database")
	}

	autoMigrate()
}

func autoMigrate() {
	// Auto Migrate Schema
	err := DB.AutoMigrate(
		&model.City{},
		&model.BusinessType{},
		&model.Location{},
		&model.Company{},
		&model.User{},
	)

	if err != nil {
		fmt.Printf("config.autoMigrate, error auto migrate: %v", err)
	}
}
