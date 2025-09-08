package handlers
import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/josephgabrie/geo-friends/repositories"
	"log"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func AddUser(db *sql.DB, c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	
	if err := repositories.InsertUsers(db, user.Username, user.Password, user.Email, user.Age); err != nil{
		return c.Status(500).JSON(fiber.Map{"error": "Could not add user"})
	}
		
	return c.SendString("Successfully Added User")
	}

func GetUserLogin(db *sql.DB, c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	dbUserName, dbEmail, dbPassword, err := repositories.CheckUserInfo(db, user.Username, user.Email)
	fmt.Print("user username" + user.Username + "\n")
	fmt.Print("user email" + user.Email + "\n")

	fmt.Print("database username" + dbUserName + "\n")
	fmt.Print("db Email" + dbEmail + "\n")

	if user.Username != dbUserName || user.Email != dbEmail {


		log.Fatal("Username or email does not match")
	}
	if user.Password != dbPassword {
		log.Fatal("Password does not match")
	}
	return err
}