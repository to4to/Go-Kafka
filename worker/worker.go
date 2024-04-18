package main

import "github.com/IBM/sarama"





func main(){
	topic:="comments"

	worker,err:=connectConsumer([]string{"localhost:29092"})

	if err!=nil{
panic(err)
	}


	worker.ConsumePartation(topic,0,sarama.OffsetOldest)
}