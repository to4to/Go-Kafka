package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comments", createComment)

	app.Listen(":3000")
}

func createComment(c *fiber.Ctx) error {
	cmt := new(Comment)

	if err := c.BodyParser(cmt); err != nil {



		log.Println(err)
	}
}
