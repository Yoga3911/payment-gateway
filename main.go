package main

import (
	"log"
	"os"
	"payment/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	app := fiber.New()
	routes.Data(app)
	log.Panic(app.Listen(":" + os.Getenv("PORT")))
}
