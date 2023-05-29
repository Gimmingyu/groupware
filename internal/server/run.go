package server

import (
	"flag"
	"groupware/internal/config"
	"log"
)

var (
	AppMode = flag.String("mode", "dev", "Application mode")
)

func Run() {

	flag.Parse()

	log.Println(AppMode)

	db := config.Database(*AppMode)
	config.Migrate(db)
}
