package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"golang.org/x/text/cases"
)





func main(){
	topic:="comments"

	worker,err:=connectConsumer([]string{"localhost:29092"})

	if err!=nil{ƒ
panic(err)
	}


consumer,err:=	worker.ConsumePartation(topic,0,sarama.OffsetOldest)

if err!=nil{
	panic(err)
		}

		fmt.Println("Consumer Started")
		sigchan:=make(chan os.Signal,1)
		signal.Notify(sigchan,syscall.SIGINT,syscall.SIGTERM)

		msgCount:=0
		doneCh:=make(chan struct{})


		go func(){

			for{
				select{
				case err:=<- consumer.Error():
					fmt.Println(err)
				case msg:=<-consumer.Messages():
					msgCount++
					fmt.Println("Received Message : %d | Topic : %s | Message(%s)n", msgCount,string(msg.Topic),string(msg.Value))
				case <-sigchan:
					fmt.Println("Interuption Detected ")
					doneCh <-struct{}{}
				}

			}
		}()
}