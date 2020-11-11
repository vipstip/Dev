package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func homePagesss(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Fprintf(w, GlolName)
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePagesss)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

var GlolName = "Tung Dep Trai"

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	go Consumer(consumer)

	handleRequests()

}

func Consumer(consumer sarama.Consumer) {
	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Println("Consumed message offset ", msg.Offset)
			log.Println("Consumed message value ", string(msg.Value))
			GlolName = string(msg.Value)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)

}
