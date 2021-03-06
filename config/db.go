package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"pos_api/model"
)

var DB *gorm.DB

func InitDb() {
	// Configure DB Connection
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		Config.DbHost, Config.DbUser, Config.DbPassword, Config.DbName, Config.DbPort, Config.DbSslMode, Config.DbTimeZone)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
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
		&model.Outlet{},
		&model.UnitOfMeasurement{},
		&model.ItemGroup{},
		&model.ItemCategory{},
		&model.Item{},
		&model.ItemVariant{},
		&model.Inventory{},
		&model.UserRole{},
		&model.Module{},
		&model.SubModule{},
	)

	if err != nil {
		log.Fatalf("config.autoMigrate, error auto migrate: %v", err)
	}

}
