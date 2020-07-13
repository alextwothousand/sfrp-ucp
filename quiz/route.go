package quiz

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/sanfierrorp/ucp/models"
	"github.com/sanfierrorp/ucp/storage"
)

// Get returns the player's quiz questions— if they have any.
func Get(c *fiber.Ctx) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"]

	db := storage.Instance()
	count := 0

	db.Where("playersqlid = ?", id).Find(&models.Quiz{}).Count(&count)

	if count < 1 {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": "Doesnt have any quiz questions.",
		})
		return
	}

	var quiz *models.Quiz
	db.Where("playersqlid = ?", id).First(&quiz)
}

// Post creates players quiz questions since they don't have any.
func Post(c *fiber.Ctx) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))

	db := storage.Instance()
	count := 0

	db.Where(&models.Quiz{PlayerSqlID: id}).Find(&models.Quiz{}).Count(&count)

	if count > 0 {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":    false,
			"error": "Already has quiz questions.",
		})
		return
	}

	questions := make(Questions, 0)

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		question := quizQuestions[rand.Intn(len(quizQuestions))]

		for _, q := range questions {
			for q.Question == question.Question {
				question = quizQuestions[rand.Intn(len(quizQuestions))]
			}
		}

		questions = append(questions, question)
	}

	/*db.Create(&models.Quiz{
		PlayerSqlID: id,
	})*/

	fmt.Println(questions)
	c.Status(200)
}
