package main

import (
	"log"

	"github.com/Valdeci97/full-cycle-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	producer := kafka.NewKafkaProducer()
	kafka.Publish("olá kafka", "readtest", producer)

	for {
		_ = 1
	}

	// r := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// r.LoadPositions()
	// stringjson, _ := r.ExportJsonPositions()
	// fmt.Println(stringjson[0])
}
