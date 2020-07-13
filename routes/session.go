package routes

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/sanfierrorp/ucp/models"
	"github.com/sanfierrorp/ucp/storage"
)

type newSession struct {
	Username string `json:"username"`
}

// NewSession creates a new session with the server.
func NewSession(c *fiber.Ctx) {
	// /api/v1/session/new - Method: GET
	c.Accepts("application/json")
	s := new(newSession)

	if err := c.BodyParser(s); err != nil {
		log.Fatal(err)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = s.Username
	claims["id"] = -1
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	db := storage.Instance()
	count := 0

	db.Where("name = ?", s.Username).Find(&models.Players{}).Count(&count)

	if count < 1 {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": "An error has occured.",
		})
		return
	}

	player := models.Players{}
	db.Where("name = ?", s.Username).First(&player)

	claims["id"] = player.ID

	t, err := token.SignedString([]byte("Bv&06I6Nun&r!#c6ra$r"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return
	}

	//fmt.Println("Token is", t)

	c.Status(200).JSON(fiber.Map{
		"token": t,
	})
	return
}

// AuthSession authenticates the users session.
func AuthSession(c *fiber.Ctx) {
	// /api/v1/session - Method: GET
	//fmt.Println("status is 200")
	return
}
