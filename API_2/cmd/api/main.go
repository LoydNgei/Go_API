package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/avukadin/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

// Chi router, logrus for logging, custom handlers from an internal package

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println("Loving Golang already")

	err := http.listenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}