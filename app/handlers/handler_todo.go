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
	// Recuperando a sessão do usuário
	sess, err := h.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := sess.Get("user_id").(int)

	// Obter tarefas do banco de dados
	rows, err := h.DB.Query("SELECT * FROM ToDos WHERE userid = $1", userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	todos := []models.ToDo{}
	for rows.Next() {
		todo := models.ToDo{}
		err := rows.Scan(&todo.UserId, &todo.Id, &todo.Title, &todo.Completed)
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
	// recuperar a sessão do usuário
	sess, err := h.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Obter o ID do usuário da sessão
	userID := sess.Get("user_id").(int)

	// Obter os dados da tarefa do corpo da requisição
	todo := &models.ToDo{}
	err = c.BodyParser(todo)
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

	// Verificar se a tarefa existe no banco de dados e se pertence ao usuário
	row := h.DB.QueryRow("SELECT id FROM ToDos WHERE id = $1 and userid = $2", todo.Id, userID)
	err = row.Scan(&todo.Id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Tarefa não encontrada",
		})
	}

	// Construir dinamicamente a consulta SQL UPDATE
	query := "UPDATE ToDos SET "
	args := []interface{}{}
	cont := 1
	if todo.Title != "" {
		query += fmt.Sprintf("title = $%d, ", cont)
		args = append(args, todo.Title)
		cont++
	}
	if todo.Completed || !todo.Completed {
		query += fmt.Sprintf("completed = $%d, ", cont)
		args = append(args, todo.Completed)
		cont++
	}
	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" WHERE id = $%d and userid = $%d", cont, cont+1)
	args = append(args, todo.Id)
	args = append(args, userID)

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

func (h *Handlers) DeleteTarefa(c *fiber.Ctx) error {
	//Obter o ID da tarefa da URL
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Verificar se a tarefa existe no banco de dados
	row := h.DB.QueryRow("SELECT id FROM ToDos WHERE id = $1", id)
	err = row.Scan(&id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Tarefa não encontrada",
		})
	}

	// Excluir tarefa do banco de dados
	_, err = h.DB.Exec("DELETE FROM ToDos WHERE id = $1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Tarefa excluída com sucesso",
	})
}
