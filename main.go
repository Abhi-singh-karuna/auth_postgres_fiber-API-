package main

import (
	"github.com/abhishek_singh/database"
	"github.com/abhishek_singh/router"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// initialise the fiber
	app := fiber.New()

	//initialize the logger
	app.Use(logger.New())

	//load the .env file
	godotenv.Load()

	//connect the database
	database.ConnectDB()
	defer database.DB.Close()

	//call the routes folder
	router.Router(app)

	//listen on this port
	app.Listen(":4000")

}
