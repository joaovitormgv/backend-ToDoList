package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// AuthRequired é um middleware que verifica se o usuário está autenticado
func AuthRequired(Store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if sess.Get("user_id") == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Não autorizado",
			})
		}

		return c.Next()
	}
}
