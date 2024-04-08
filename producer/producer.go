package main

import (
	"encoding/json"
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

func ConnectProducer(brokersUrl string)(srama.SyncProducer,error){
	config:=sarama.NewConfig()
	config.Pr
}



func PushCommentToQueue(topic string,message []byte){
brokersUrl:=[]string{"localhost:29092"}

producer ,err:=ConnectProducer(brokersUrl)



}

func createComment(c *fiber.Ctx) error {
	cmt := new(Comment)

	if err := c.BodyParser(cmt); err != nil {

		log.Println(err)

		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}

	cmtInBytes, err := json.Marshal(cmt)

	PushCommentToQueue("comments",cmtInBytes)

	err=c.JSON(&fiber.Map{
		"success":true,
		"message":"Comment pushed Successfully",
		"comment":cmt,
	})


	if err!=nil{
		c.Status(500).JSON(&fiber.Map{
			"success":false,
			"message":"Error creating product",
		})
		return err
	}
	return err
}
