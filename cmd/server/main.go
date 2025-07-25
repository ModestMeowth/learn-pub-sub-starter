package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ModestMeowth/learn-pub-sub-starter/internal/pubsub"
	"github.com/ModestMeowth/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

const serverUrl = "amqp://guest:guest@localhost:5672"

func main() {
	fmt.Println("Starting Peril server...")
    conn, err := amqp.Dial(serverUrl)
    if err != nil {
        log.Fatalf("error: %v\n", err.Error())
    }
    defer conn.Close()

    fmt.Printf("Connection to %s successful\n", serverUrl)

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, os.Interrupt)
    <-signalChan

    c, err := conn.Channel()
    if err != nil {
        log.Fatalf("error: %v\n", err.Error())
    }
    defer c.Close()

    err = pubsub.PublisJSON(c, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{
        IsPaused: true,
    })

    fmt.Printf("\nInterrupt signal received, shutting down\n")
}
