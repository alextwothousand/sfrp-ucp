package routes

import (
	"github.com/gofiber/fiber"
	"github.com/sanfierrorp/ucp/models"
	"github.com/sanfierrorp/ucp/storage"
)

// News retrieves news articles from the database.
func News(c *fiber.Ctx) {
	db := storage.Instance()

	var news []models.News
	db.Find(&news)

	c.JSON(news)
	return
}
