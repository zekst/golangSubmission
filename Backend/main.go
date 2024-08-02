package main

import (
	"Module/database"
	"Module/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "http://localhost, http://localhost:3000/",
	}))
	
	routes.Setup(app)

	app.Listen(":8000")
}