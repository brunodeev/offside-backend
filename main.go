package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Lista de produtos aqui! Este é um teste do pipeline!"))
	})

	app.Listen(":8080")
}
