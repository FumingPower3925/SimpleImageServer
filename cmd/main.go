package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/FumingPower3925/SimpleImageServer/pkg/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")

	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	go serv.Start()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt)
	<-sign

	serv.Close()
}
