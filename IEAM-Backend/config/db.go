package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ieam-backend/models"
)

var DB *gorm.DB

func ConnectDB() error {
	var err error
	var dsn = "host=localhost user=postgres dbname=IEAM port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("postgres err: " + err.Error())
		panic(err)
	}
	migrateDB()
	return err
}

func migrateDB() {
	DB.AutoMigrate(models.Account{})
	DB.AutoMigrate(models.Form{})
	DB.AutoMigrate(models.List_Form{})
	// DB.AutoMigrate(models.History{})
}