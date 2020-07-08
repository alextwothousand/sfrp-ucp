package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber"
	"github.com/sanfierrorp/ucp/routes"
	"github.com/sanfierrorp/ucp/storage"
)

func main() {
	storage.Open()
	app := fiber.New()

	app.Get("/api/news", routes.News)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)

	go func() {
		<-sc
		storage.Close()
		fmt.Println("Close database, or better create a func here")
		app.Shutdown()
	}()

	app.Listen(3001)
}
