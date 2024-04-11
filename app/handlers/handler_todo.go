package handlers

import (
	"fmt"
	"strings"

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

func (h *Handlers) GetTarefas(c *fiber.Ctx) error {
	// Obter tarefas do banco de dados
	rows, err := h.DB.Query("SELECT * FROM ToDos")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	todos := []models.ToDo{}
	for rows.Next() {
		todo := models.ToDo{}
		err := rows.Scan(&todo.Id, &todo.UserId, &todo.Title, &todo.Completed)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func (h *Handlers) UpdateTarefa(c *fiber.Ctx) error {
	todo := &models.ToDo{}
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//Obter o ID da tarefa da URL
	todo.Id, err = c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validar os dados da tarefa
	//...

	// Construir dinamicamente a consulta SQL UPDATE
	query := "UPDATE ToDos SET "
	args := []interface{}{}
	cont := 1
	if todo.Title != "" {
		query += fmt.Sprintf("title = $%d, ", cont)
		args = append(args, todo.Title)
		cont++
	}
	if todo.UserId != 0 {
		query += fmt.Sprintf("userid = $%d, ", cont)
		args = append(args, todo.UserId)
		cont++
	}
	if todo.Completed || !todo.Completed {
		query += fmt.Sprintf("completed = $%d, ", cont)
		args = append(args, todo.Completed)
		cont++
	}
	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" WHERE id = $%d", cont)
	args = append(args, todo.Id)

	// Atualizar tarefa no banco de dados
	_, err = h.DB.Exec(query, args...)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Tarefa atualizada com sucesso",
	})
}
