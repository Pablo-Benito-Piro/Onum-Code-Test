package models

type Auction struct {
	ID           string `json:"id" gorm:"primary_key;unique;column:AuctionID;"`
	StartValue   int64  `json:"startValue" gorm:"column:start_value;"`
	BidStartTime int64  `json:"bidStartTime" gorm:"column:bid_start_time;type:bigint(20);not null"`
	BidEndTime   int64  `json:"bidEndTime" gorm:"column:bid_end_time;type:bigint(20);not null"`
}

type AuctionSearch struct {
	BidStartTime int64 `json:"bidStartTime"`
	BidEndTime   int64 `json:"bidEndTime" `
}

const (
	Processing string = "processing"
	Best       string = "best"
	Outbided   string = "outbided"
	Won        string = "won"
	Lost       string = "lost"
)
