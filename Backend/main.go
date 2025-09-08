package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/josephgabrie/geo-friends/handlers"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in .env")
	}

	var connErr error
	db, connErr = sql.Open("postgres", dsn)
	if connErr != nil {
		log.Fatal(connErr)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Cannot connect to database:", pingErr)
	}
	log.Println("Connected to database successfully!")
}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!!!")
	})
	app.Post("/friends", func(c *fiber.Ctx) error {
		return handlers.AddFriend(db, c)
	})
	app.Get("/friend", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	app.Post("/userSignUp", func(c *fiber.Ctx) error {
		return handlers.AddUser(db, c)
	})
	app.Get("/userSignIn", func(c *fiber.Ctx) error {
		return handlers.GetUserLogin(db, c)
	})
	app.Listen(":3000")
}
