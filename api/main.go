package main

import (
	"fmt"
	"log"
	"os"
	"github.com/Genialngash/golang-url-shortener/api/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)



func setUpRoutes(app *fiber.App){
	app.Get("/:url",routes.ResolveURL)
	app.Post("/api/v1", routes.ShortedURL)

}


func main(){
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	
	app := fiber.New()

	app.Use(logger.New())
	setUpRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}