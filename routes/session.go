package routes

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
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
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("Bv&06I6Nun&r!#c6ra$r"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return
	}

	//fmt.Println("Token is", t)

	c.JSON(fiber.Map{
		"token": t,
	})
	c.Status(200)
	return
}

// AuthSession authenticates the users session.
func AuthSession(c *fiber.Ctx) {
	// /api/v1/session - Method: GET

}

// DeleteSession clears the users session.
func DeleteSession(c *fiber.Ctx) {
	// /api/v1/session - Method: DELETE
}
