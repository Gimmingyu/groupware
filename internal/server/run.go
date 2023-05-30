package server

import (
	"flag"
	"groupware/internal/config"
)

var (
	AppMode = flag.String("mode", "development", "Application mode")
)

func Run() {

	flag.Parse()

	db := config.Database(*AppMode)
	config.Migrate(db)

}