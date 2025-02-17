package dao

import (
	"onumTest/commons"
	"onumTest/models"
)

func CreateBidDAO(bid models.Bid) (models.Bid, error) {
	db := commons.GetConnection()
	err := db.Create(&bid).Error
	db.Close()
	return bid, err
}

func FindBidsByAuctionIDAndClientID(clientID string, auctionID string) []models.Bid {
	var bids []models.Bid

	db := commons.GetConnection()
	defer db.Close()
	db.Find(&bids, "AuctionID = ? and clientId = ?", auctionID, clientID)
	return bids
}
func FindBidsByStatusBest(auctionID string) models.Bid {
	var bid models.Bid
	db := commons.GetConnection()
	defer db.Close()
	db.Find(&bid, "status = '"+models.Best+"' and AuctionID = ?", auctionID)
	return bid
}

func FindBids(auctionID string) []models.Bid {
	var bid []models.Bid
	db := commons.GetConnection()
	db.Find(&bid, "AuctionID = ?", auctionID)
	db.Close()
	return bid
}
func SaveBid(bid models.Bid) {
	db := commons.GetConnection()
	db.Model(bid).Where("id = ?", bid.ID).Update("status", bid.Status)
	db.Close()

}
