package main

import (
	"log"

	"fst/project/database"
	"fst/project/routes"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to an Awesome API")
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)

	// User endpoints
	app.Post("/api/signup", routes.SignUp)
	app.Post("/api/login", routes.Login)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/user/:id", routes.GetUser)
	app.Delete("/api/user/:id", routes.DeleteUser)
	// Heartbeat Socket
	app.Get("/ws/heartbeat/:id", websocket.New(routes.Heartbeat_socket))
	app.Get("/ws/sleep/:id", websocket.New(routes.Sleep_socket))
	app.Get("/ws/step/:id", websocket.New(routes.Step_socket))
	app.Post("api/heartbeat", routes.CreateHeartbeat)

}

func main() {

	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDb()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
