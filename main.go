package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joaovitormgv/backend-ToDoList/app/handlers"
	"github.com/joaovitormgv/backend-ToDoList/app/middleware"
	"github.com/joaovitormgv/backend-ToDoList/app/routes"
)

func main() {
	user := "postgres"
	password := "123456"
	dbname := "todolist"
	connectionString := fmt.Sprintf("postgres://%s:%s@localhost:5433/%s?sslmode=disable", user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := session.New()

	h := &handlers.Handlers{
		DB:    db,
		Store: store,
	}
	app := fiber.New()
	app.Use(middleware.CorsMiddleware())
	routes.Setup(app, h)
	app.Listen(":3000")
}
