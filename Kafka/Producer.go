package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main()  {
	for {
		config := sarama.NewConfig()
		config.Producer.Return.Successes = true
		producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
		if err != nil {
			panic(err)
		}
		fmt.Print("Enter text: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Println("++++",strings.TrimRight(text, "\n"),"++++")
		if strings.TrimRight(text, "\r\n") == "exit"{
			break
		}
		//Trap SIGINT to trigger a graceful shutdown.
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		message := &sarama.ProducerMessage{Topic: "ok", Value: sarama.StringEncoder(text)}
		select {
		case producer.Input() <- message:
			log.Printf("Successfully produced: ",text)
		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break
		}
	}
}