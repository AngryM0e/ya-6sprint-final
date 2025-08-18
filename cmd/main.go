package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func main() {
	if err := handlers.LoadTemplate(); err != nil {
		log.Fatal("Не удалось загрузить index.html:", err)
	}

	logger := log.New(os.Stdout, "morse: ", log.LstdFlags|log.Lshortfile)
	srv := server.New(logger)

	if err := srv.Start(); err != nil {
		logger.Fatal(err)
	}
}
