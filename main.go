package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joaovitormgv/backend-ToDoList/app/handlers"
	"github.com/joaovitormgv/backend-ToDoList/app/routes"
)

func main() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	connectionString := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbname)

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
	routes.Setup(app, h)
	app.Listen(":3000")
}
