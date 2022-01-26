package route

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/learn-go/controllers/public"
)

func InitRouter() {
	app := fiber.New()

	app.Get("/public/getTimestamp", public.GetTimestamp)

	log.Fatal(app.Listen(":3000"))
}
