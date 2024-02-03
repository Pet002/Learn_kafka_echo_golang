package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"demo.kafka.database/app"
	"demo.kafka.database/app/consumer"
	"demo.kafka.database/app/productor"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq" // The database driver in use.
)

func main() {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", "localhost", "5432", "postgres", "password", "user"))
	if err != nil {
		log.Fatal("DB error")
	}

	// New Service for Kafka
	consumerStorage := consumer.NewConsumerStorage(db)
	consumerService := consumer.NewConsumerService(*consumerStorage)

	// start consumer to get data from kafka and print to log
	go app.StartConsumer(consumerService.KafKaUserTopic)

	// start producer to product a data package to kafka
	p, err := app.StartProductor()
	if err != nil {
		log.Fatal("Kafka error:" + err.Error())
	}
	defer p.Close()

	producterHandler := productor.NewProducerHandler(p)

	e := echo.New()

	e.POST("/", producterHandler.ProductorHandler)

	e.Start(":8080")
	// create signal recive to quit from server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	e.Shutdown(context.Background())
}
