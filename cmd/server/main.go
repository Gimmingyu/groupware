package main

import (
	"groupware/internal/server"
	"groupware/pkg/safety"
	"log"
	"os"
	"os/signal"
)

func main() {
	go safety.SafeBootstrap(server.Run)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	log.Println("Shutting down...")
}
