package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/backend-ToDoList/app/handlers"
)

func Setup(app *fiber.App, h *handlers.Handlers) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/cadastro/usuario", h.CreateUser)
}
