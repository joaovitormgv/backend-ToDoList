package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/backend-ToDoList/app/handlers"
	"github.com/joaovitormgv/backend-ToDoList/app/routes"
)

func main() {
	connectionString := "host=localhost port=5432 user=postgres password=123456 dbname=todolist sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := &handlers.Handlers{DB: db}
	app := fiber.New()
	routes.Setup(app, h)
	app.Listen(":3000")
}
