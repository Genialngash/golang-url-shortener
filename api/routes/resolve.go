package routes

import (
	"github.com/Genialngash/golang-url-shortener/api/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber"
)

func ResolveURL(c *fiber.Ctx) {
	url := c.Params("url")
	r := database.CreateClient(0)
	// close the db connection
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err != redis.Nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short not found in the database"})
		return 

	} else if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not connect to the db"})
		return 

	}
	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	c.Redirect(value, 301)
	
}
