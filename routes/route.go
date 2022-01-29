package route

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/learn-go/controllers/identity"
	"github.com/hoanggggg5/learn-go/controllers/public"
	resources "github.com/hoanggggg5/learn-go/controllers/resource"
)

func InitRouter() {
	app := fiber.New()

	app.Get("/public/getTimestamp", public.GetTimestamp)
	app.Post("/resource/users/create", resources.CreateUser)
	app.Post("/login", identity.Login)
	app.Get("/getme", resources.GetMe)

	log.Fatal(app.Listen(":3000"))
}
