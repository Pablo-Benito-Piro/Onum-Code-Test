package models

type Bid struct {
	BidID     string  `json:"bidId" gorm:"primary_key;unique;column:BidID;"`
	AuctionID string  `json:"auctionID" gorm:"column:AuctionID;"`
	Auction   Auction `gorm:"foreignKey:AuctionID;references:AuctionID"`
	Bid       int64   `json:"bid" gorm:"column:bid;"`
	Status    string  `json:"status" gorm:"column:status;"`
	ClientId  int64   `json:"clientId" gorm:"column:clientId;"`
}
