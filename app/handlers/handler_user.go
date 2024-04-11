package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/backend-ToDoList/app/models"
)

type Handlers struct {
	DB *sql.DB
}

func (h *Handlers) CreateUser(c *fiber.Ctx) error {
	user := &models.User{}
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validar os dados do usuário
	if user.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Nome é obrigatório",
		})
	}

	// Inserir usuário no banco de dados
	row := h.DB.QueryRow("INSERT INTO usuario (nome) VALUES ($1) RETURNING id", user.Name)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":   id,
		"name": user.Name,
	})
}
