package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"tanachod.learn-project/handler"
)

func main() {
	e := echo.New()

	userHandler := handler.NewUserHandler(handler.User{
		Name:    "Tanachod",
		Surname: "Sakthamjaroen",
	})
	e.GET("/", userHandler.GetHandler)

	go startEcho(e)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	e.Logger.Printf(" go routine will running")
	<-quit
	e.Logger.Printf(" go routine will ending")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func startEcho(e *echo.Echo) {
	if err := e.Start(":8080"); err != nil {
		log.Print(err)
	}

}
