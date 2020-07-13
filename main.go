package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/sanfierrorp/ucp/quiz"
	"github.com/sanfierrorp/ucp/routes"
	"github.com/sanfierrorp/ucp/storage"

	jwtware "github.com/gofiber/jwt"
)

func main() {
	storage.Open()
	app := fiber.New()

	// Allows access from Cross Origin regions. In our case, port 3000 is our next.js application.
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPatch,
		},
	}))

	// Handle news articles.
	app.Get("/api/v1/news", routes.News) // Retrieve news articles.
	//app.Post("/api/v1/news", routes.AddNews) / / Create a new news article.

	// Handle users.
	app.Post("/api/v1/user", routes.PostUser) // Register a brand new user.
	app.Get("/api/v1/user", routes.GetUser)   // Get a user's general information.

	// Handle sessions.
	app.Post("/api/v1/session", routes.NewSession) // Creates a new session.

	// JSON Web Tokens middleware. Allows user session management to occur! :)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("Bv&06I6Nun&r!#c6ra$r"),
	}))

	// Handle users.
	//app.Get("/api/v1/user/:id", routes.GetUser)

	// Handle sessions.
	app.Get("/api/v1/session", routes.AuthSession) // Authenticates a user's session.

	// Handle Quiz Questions
	app.Get("/api/v1/quiz", quiz.Get)
	app.Post("/api/v1/quiz", quiz.Post)

	// Handle characters.
	//app.Get("/api/v1/characters")
	//app.Get("/api/v1/characters/:id")

	//app.Post("/api/v1/characters")
	//app.Patch("/api/v1/characters/:id")
	//app.Delete("/api/v1/characters/:id")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)

	// if OS INTERRUPT is received, the database will be closed and so will the go-fiber application.
	go func() {
		<-sc
		storage.Close()
		app.Shutdown()
	}()

	// Listen on port 3001.
	app.Listen(3001)
}
