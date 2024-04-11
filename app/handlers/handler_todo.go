package handlers

import (
	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/backend-ToDoList/app/models"
)

func (h *Handlers) CreateTarefa(c *fiber.Ctx) error {
	todo := &models.ToDo{}
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validar os dados da tarefa
	if todo.UserId <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "UserId é obrigatório",
		})
	} else if todo.Title == "" {
		todo.Title = "Tarefa sem título"
	}

	// Inserir tarefa no banco de dados
	row := h.DB.QueryRow("INSERT INTO ToDos (userid, title, completed) VALUES ($1, $2, $3) RETURNING id", todo.UserId, todo.Title, "false")
	var id int
	err = row.Scan(&id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":        id,
		"userId":    todo.UserId,
		"title":     todo.Title,
		"completed": "false",
	})
}
