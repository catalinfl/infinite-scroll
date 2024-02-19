package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"

	"github.com/catalinfl/infinite-scroll/handlers"
	"github.com/catalinfl/infinite-scroll/utils"
)

func main() {

	app := fiber.New()

	utils.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173",
	}))

	app.Get("/api", handlers.PaginationUser)
	app.Get("/api", handlers.GetUsers)
	app.Get("/api/test", handlers.CursorPaginationPosts)
	app.Get("/api/:id", handlers.GetUser)
	app.Post("/api", handlers.CreateUser)
	app.Post("/api/post", handlers.CreateDescription)

	app.Listen(":3000")

}
