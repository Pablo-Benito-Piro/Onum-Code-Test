package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"onumTest/commons"
	"onumTest/controller"
	"onumTest/service"
	"time"
)

func main() {
	commons.InitialMigrate()
	router := mux.NewRouter()
	controller.SetAuctionRoutes(router)
	controller.SetBidRoutes(router)

	server := http.Server{
		Handler: router,
		Addr:    ":8090",
	}

	ticker := time.NewTicker(1 * time.Second) // Ejecutar cada 1 segundos
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			service.ChecksAuctionsEndTime()
		}
	}()

	log.Println(server.ListenAndServe())
}
