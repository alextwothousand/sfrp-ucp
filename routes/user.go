package routes

import (
	"fmt"
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

type loginPlayer struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	validUsername *regexp.Regexp = regexp.MustCompile(`^([A-Z]|[a-z]){3,255}$`)
	validEmail    *regexp.Regexp = regexp.MustCompile(`^.+@\w+\..+$`)
)

// PostUser handles the adduser stuff
func PostUser(c *fiber.Ctx) {
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

	if !validUsername.MatchString(p.Username) {
		c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Bad username",
		})
	}

	if !validEmail.MatchString(p.Email) {
		c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Bad email",
		})
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

// GetUser grabs a user's information IF they are registered.
func GetUser(c *fiber.Ctx) {
	c.Accepts("application/json")
	p := new(loginPlayer)

	if err := c.BodyParser(p); err != nil {
		log.Fatal(err)
	}

	if !validUsername.MatchString(p.Username) {
		c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Bad username",
		})
	}

	if !password.Integrity(p.Password) {
		c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Poor password entered",
		})
		return
	}

	db := storage.Instance()

	count := 0
	db.Where("name = ?", p.Username).Find(&models.Players{}).Count(&count)

	if count != 1 {
		c.Status(405).JSON(fiber.Map{
			"ok":    false,
			"error": "User does not exist",
		})
		return
	}

	player := models.Players{}
	db.Where("name = ?", p.Username).First(&player)

	//fmt.Println("count is", count)
	//fmt.Println("player len is", len(player))

	if !password.Match(p.Password, player.Password) {
		c.Status(401).JSON(fiber.Map{
			"ok":    false,
			"error": "Incorrect password entered",
		})

		fmt.Println("Incorrect passsword entered - tsk tsk.")
		return
	}

	c.Status(200)

	/*if p.Password != p.CPassword {
		c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Passwords do not match",
		})
		return
	}*/

	//c.Redirect("/ucp", 302)
	return
}
