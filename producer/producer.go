package main
import "github.com/gofiber/fiber/v2"

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
app:=fiber.New();
api:=app.Group("/api/v1")
}
