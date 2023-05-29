package main

import (
	"groupware/internal/server"
	"log"
	"os"
	"os/signal"
)

func main() {

	go server.Run()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	log.Println("Shutting down...")
}
