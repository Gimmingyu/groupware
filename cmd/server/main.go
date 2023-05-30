package main

import (
	"github.com/joho/godotenv"
	"groupware/internal/server"
	"log"
	"os"
	"os/signal"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	go server.Run()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	log.Println("Shutting down...")
}