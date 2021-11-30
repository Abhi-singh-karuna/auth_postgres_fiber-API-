package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abhishek_singh/model"

	"github.com/jinzhu/gorm"
)

// DB gorm connector
var DB *gorm.DB

// connection with database
func ConnectDB() {
	var err error
	p := os.Getenv("POSTGRES_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	DB, err = gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_ HOST"),
		port, os.Getenv("POSTGRES_USER_NAME"),
		os.Getenv("POSTGRES_USER_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE_NAME")))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	//for migrates the struct in the database
	DB.AutoMigrate(&model.User{}, &model.Token{})
	fmt.Println("Database Migrated")
}
