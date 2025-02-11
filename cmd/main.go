package main

import (
	"go-kafka-poc-consumer/config"
	"go-kafka-poc-consumer/internal/infra/database"
	"go-kafka-poc-consumer/internal/infra/kafka"
	"go-kafka-poc-consumer/worker"

	geutils "github.com/GenaroDaniel/geutils/pkg/events"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	db := database.Connect()
	database.Migrate(db)

	defer db.Close()
	eventDispatcher := geutils.NewEventDispatcher()
	kafkaClient, err := kafka.NewKafka(config.Config.KafkaSeeds, config.Config.KafkaTopics, *eventDispatcher)
	if err != nil {
		panic(err)
	}

	kafkaWorker := worker.NewWorker(db, kafkaClient, eventDispatcher)
	kafkaWorker.Start()

}
