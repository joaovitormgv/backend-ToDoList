package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/backend-ToDoList/app/routes"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
