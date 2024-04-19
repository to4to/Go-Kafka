package main

import (
	"fmt"

	"github.com/IBM/sarama"
)





func main(){
	topic:="comments"

	worker,err:=connectConsumer([]string{"localhost:29092"})

	if err!=nil{
panic(err)
	}


consumer,err:=	worker.ConsumePartation(topic,0,sarama.OffsetOldest)

if err!=nil{
	panic(err)
		}

		fmt.Println("Consumer Started")
}