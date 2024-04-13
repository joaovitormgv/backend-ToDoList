package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/backend-ToDoList/app/handlers"
)

func Setup(app *fiber.App, h *handlers.Handlers) {
	// Métodos para manipular usuários
	app.Post("/api/cadastro/usuario", h.CreateUser)

	// Métodos para manipular tarefas
	app.Post("/api/cadastro/tarefa", h.CreateTarefa)

	app.Get("/api/tarefas", h.GetTarefas)

	app.Put("/api/tarefa/:id", h.UpdateTarefa)

	app.Delete("/api/tarefa/:id", h.DeleteTarefa)
}
