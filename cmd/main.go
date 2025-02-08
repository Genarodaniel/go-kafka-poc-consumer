package main

import (
	"go-kafka-order-producer/config"
	"go-kafka-order-producer/internal/infra/database"
	"go-kafka-order-producer/internal/server"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	db := database.Connect()
	defer db.Close()

	database.Migrate(db)

	s := server.Init(db)
	s.Run()

}
