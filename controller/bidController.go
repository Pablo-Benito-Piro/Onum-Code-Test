package controller

import (
	"github.com/gorilla/mux"
	"onumTest/service"
)

func SetBidRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/bid").Subrouter()
	subRouter.HandleFunc("/create", service.CreateBid).Methods("POST")
	subRouter.HandleFunc("/search", service.GetBidsByAuctionIDandyClientID).Methods("GET")
	subRouter.HandleFunc("/notify", service.ClientSimulatorCallBack).Methods("PUT")

}
