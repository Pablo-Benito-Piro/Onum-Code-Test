package dao

import (
	"onumTest/commons"
	"onumTest/models"
)

func CreateAuctionDAO(auction models.Auction) error {
	db := commons.GetConnection()
	defer db.Close()
	return db.Create(&auction).Error
}

func FindAuctionByID(id string) models.Auction {
	auction := models.Auction{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&auction, "AuctionID = ?", id)

	return auction
}
func FindAuctionByStartTimeAndEndTime(startTime int64, endTime int64) []models.Auction {
	var auction []models.Auction
	db := commons.GetConnection()
	defer db.Close()
	db.Find(&auction, "bid_start_time <= ? and bid_end_time >= ?", startTime, endTime)

	return auction
}
func FindAuctions() []models.Auction {
	var auction []models.Auction
	db := commons.GetConnection()
	db.Find(&auction)
	defer db.Close()
	return auction
}
