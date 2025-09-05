package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/josephgabrie/geo-friends/repositories"
	"fmt"
)
type AddUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type AddFriendRequest struct {
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
}

type Ratings struct {
	Rating  int `json:"rating"`
	PlaceID int `json:"placeID"`
}

func AddUser(db *sql.DB, c *fiber.Ctx) error {
	var user AddUser
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	
	if err := 

}
func AddFriend(db *sql.DB, c *fiber.Ctx) error {
	var req AddFriendRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := repositories.InsertFriend(db, req.UserID, req.FriendID ); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not add friend"})
	}

	return c.JSON(fiber.Map{"Message": "Friend added succesfully"})
}
