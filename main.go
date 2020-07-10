package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/sanfierrorp/ucp/routes"
	"github.com/sanfierrorp/ucp/storage"

	jwtware "github.com/gofiber/jwt"
)

func main() {
	storage.Open()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
	}))

	app.Get("/api/v1/news", routes.News)
	app.Post("/api/v1/user", routes.AddUser)

	app.Post("/api/v1/session/new", routes.NewSession)
	app.Get("/api/v1/session", routes.AuthSession)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("Bv&06I6Nun&r!#c6ra$r"),
	}))

	app.Delete("/api/v1/session", routes.DeleteSession)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)

	go func() {
		<-sc
		storage.Close()
		app.Shutdown()
	}()

	app.Listen(3001)
}
