package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/backend-ToDoList/app/handlers"
	"github.com/joaovitormgv/backend-ToDoList/app/middleware"
)

func Setup(app *fiber.App, h *handlers.Handlers) {
	// Métodos para manipular usuários
	app.Post("/api/cadastro/usuario", h.CreateUser)

	app.Post("/api/login", h.AuthenticateUser)

	app.Post("/api/logout", middleware.AuthRequired(h.Store), h.LogoutUser)

	// Métodos para manipular tarefas
	app.Post("/api/cadastro/tarefa", middleware.AuthRequired(h.Store), h.CreateTarefa)

	app.Get("/api/tarefas", middleware.AuthRequired(h.Store), h.GetTarefas)

	app.Put("/api/tarefa/:id", middleware.AuthRequired(h.Store), h.UpdateTarefa)

	app.Delete("/api/tarefa/:id", middleware.AuthRequired(h.Store), h.DeleteTarefa)
}
