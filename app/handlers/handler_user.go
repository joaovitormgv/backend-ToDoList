package handlers

import (
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/backend-ToDoList/app/models"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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
	if user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email é obrigatório",
		})
	} else if !strings.Contains(user.Email, "@") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email inválido",
		})
	}
	if user.Password == "" || len(user.Password) < 4 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Senha é obrigatória e deve ter no mínimo 4 caracteres",
		})
	}

	// Verificar se o email já está cadastrado
	row := h.DB.QueryRow("SELECT email FROM usuario WHERE email = $1", user.Email)
	var email string
	err = row.Scan(&email)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email já cadastrado",
		})
	}

	// encriptar senha
	password := user.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user.Password = string(hashedPassword)

	// Inserir usuário no banco de dados
	row = h.DB.QueryRow("INSERT INTO usuario (nome, email, senha) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Email, user.Password)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":    id,
		"name":  user.Name,
		"email": user.Email,
	})
}
