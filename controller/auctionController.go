package controller

import (
	"github.com/gorilla/mux"
	"onumTest/service"
)

func SetAuctionRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/auction").Subrouter()
	subRouter.HandleFunc("/create", service.CreateAuction).Methods("POST")
	subRouter.HandleFunc("", service.GetAuctionsByStartTimeAndEndTime).Methods("GET")
}
