package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/sanfierrorp/ucp/models"
	"github.com/sanfierrorp/ucp/storage"
	"golang.org/x/crypto/bcrypt"
)

type player struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPassword string `json:"cPassword"`
}

// AddUser handles the adduser stuff
func AddUser(c *fiber.Ctx) {
	c.Accepts("application/json")
	p := new(player)

	if err := c.BodyParser(p); err != nil {
		log.Fatal(err)
	}

	if len(p.Username) < 1 || len(p.Password) < 1 || len(p.CPassword) < 1 || len(p.Email) < 1 {
		fmt.Println(c.Body())
		c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Incorrect input lengths",
		})
		return
	}

	if p.Password != p.CPassword {
		c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Passwords do not match",
		})
		return
	}

	db := storage.Instance()

	pass, err := bcrypt.GenerateFromPassword([]byte(p.Password), 12)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return
	}

	db.Create(&models.Players{
		Name:       p.Username,
		Password:   string(pass),
		Email:      p.Email,
		RegisterIP: "NULL",
		IP:         "NULL",
		Donator:    0,
		Helper:     0,
		AdminLevel: 0,
	})

	c.Status(200)
	return
}
