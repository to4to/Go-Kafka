package main

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
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

func ConnectProducer(brokersUrl string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successs = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil

}
func PushCommentToQueue(topic string, message []byte) {
	brokersUrl := []string{"localhost:29092"}
	producer, err := ConnectProducer(brokersUrl[0])

if err!=nil{

}
defer producer.Close()
msg:=&sarama.ProducerMessage{
	Topic:topic,
	Value:sarama.StringEncoder(message),
}

partation,offset,err:= producer.SendMessage(msg)

if err !=nil{}
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

	PushCommentToQueue("comments", cmtInBytes)

	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "Comment pushed Successfully",
		"comment": cmt,
	})

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return err
	}
	return err
}
