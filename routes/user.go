package routes

import (
	"log"
	"regexp"

	"github.com/gofiber/fiber"
	"github.com/sanfierrorp/ucp/models"
	"github.com/sanfierrorp/ucp/password"
	"github.com/sanfierrorp/ucp/storage"
)

type player struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPassword string `json:"cPassword"`
}

var (
	validUsername *regexp.Regexp = regexp.MustCompile(`^([A-Z]|[a-z])+$`)
	validEmail    *regexp.Regexp = regexp.MustCompile(`^.+@\w+\..+$`)
)

// AddUser handles the adduser stuff
func AddUser(c *fiber.Ctx) {
	c.Accepts("application/json")
	p := new(player)

	if err := c.BodyParser(p); err != nil {
		log.Fatal(err)
	}

	if !password.Integrity(p.Password) || !password.Integrity(p.CPassword) {
		c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Poor password entered",
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

	pass, err := password.Hash(p.Password)
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
