package db

import (
	"fmt"

	"github.com/xdarkyne/steamgo/config"
	"github.com/xdarkyne/steamgo/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ORM *gorm.DB

func Connect() {
	var (
		err error
		dsn string
		cfg *gorm.Config
	)

	cfg = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.App.DB.Host,
		config.App.DB.Username,
		config.App.DB.Password,
		config.App.DB.DBName,
		config.App.DB.Port,
		config.App.DB.SSLMode,
		config.App.TimeZone,
	)

	ORM, err = gorm.Open(postgres.Open(dsn), cfg)
	ORM.Session(&gorm.Session{FullSaveAssociations: true})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("Connected to database successfully")
}

func Migrate() {
	fmt.Println("Running migrations")

	err := ORM.Migrator().AutoMigrate(
		&models.Session{},
		&models.Category{},
		&models.Genre{},
		&models.Game{},
		&models.User{},
	)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("Completed migrations")
}
