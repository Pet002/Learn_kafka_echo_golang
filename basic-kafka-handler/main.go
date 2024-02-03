package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"demo.kafka/app"
	"demo.kafka/app/productor"
	"github.com/labstack/echo/v4"
)

func main() {

	// start consumer to get data from kafka and print to log
	go app.StartConsumer()

	// start producer to product a data package to kafka
	p, err := app.StartProductor()
	if err != nil {
		log.Fatal("Kafka error:" + err.Error())
	}
	defer p.Close()

	producterHandler := productor.NewProducerHandler(p)

	e := echo.New()

	e.GET("/", producterHandler.ProductorHandler)

	e.Start(":8080")
	// create signal recive to quit from server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	e.Shutdown(context.Background())
}
