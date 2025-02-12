package models

type Bid struct {
	ID        uint   `json:"bidId" gorm:"primary_key;unique;auto_increment:true"`
	AuctionID string `json:"auctionID" gorm:"column:AuctionID;"`
	Bid       int64  `json:"bid" gorm:"column:bid;"`
	Status    string `json:"status" gorm:"column:status;"`
	ClientId  string `json:"clientId" gorm:"column:clientId;"`
	Update    string `json:"update" gorm:"column:update;"`
}

type BidSearch struct {
	ClientID  string `json:"clientID"`
	AuctionID string `json:"auctionID" `
}
