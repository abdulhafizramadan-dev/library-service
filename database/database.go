package database

import (
	"fmt"
	"gojek.com/abdul/prebootcamp/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func Connect() error {
	portString := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portString)

	if err != nil {
		fmt.Println("DB_PORT is not a number")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	db.AutoMigrate(&model.Book{})

	DB = DbInstance{
		Db: db,
	}

	return nil
}
