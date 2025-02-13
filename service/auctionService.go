package service

import (
	"encoding/json"
	"io"
	"net/http"
	"onumTest/commons"
	"onumTest/dao"
	"onumTest/models"
	"time"
)

func GetAuctionsByStartTimeAndEndTime(w http.ResponseWriter, r *http.Request) {
	var auctions []models.Auction
	auctionSearch := models.AuctionSearch{}
	body, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal(body, &auctionSearch); err != nil {
		commons.SendError([]byte(err.Error()), w, http.StatusNotFound)
		return
	}

	db := commons.GetConnection()
	defer db.Close()

	auctions = dao.FindAuctionByStartTimeAndEndTime(auctionSearch.BidStartTime, auctionSearch.BidEndTime)

	if len(auctions) == 0 {
		commons.SendAuctionEndDateError(w, http.StatusNotFound)
	} else {
		result, _ := json.Marshal(auctions)
		commons.SendResponse(w, http.StatusOK, result)
	}

}
func ChecksAuctionsEndTime() {

	auctions := dao.FindAuctions()
	for _, auction := range auctions {
		if auction.BidEndTime < time.Now().UnixMilli() {
			bids := dao.FindBids(auction.ID)
			for _, bid := range bids {
				if bid.Status == models.Best || bid.Status == models.Won {
					bid.Status = models.Won
				} else {
					bid.Status = models.Lost
				}
				dao.SaveBid(bid)
				notifyClient(bid)
			}

		}
	}
}

func CreateAuction(w http.ResponseWriter, r *http.Request) {
	auction := models.Auction{}

	error := json.NewDecoder(r.Body).Decode(&auction)

	if error != nil {
		commons.SendError([]byte(error.Error()), w, http.StatusBadRequest)
		return
	}

	if auction.BidEndTime < time.Now().UnixMilli() {
		commons.SendAuctionEndDateError(w, http.StatusBadRequest)
		return
	}
	result := dao.FindAuctionByID(auction.ID)

	if result.ID == auction.ID {
		commons.SendADuplicateAuctionError(w, http.StatusInternalServerError)
		return
	}

	error = dao.CreateAuctionDAO(auction)

	if error != nil {
		commons.SendError([]byte(error.Error()), w, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(auction)

	commons.SendResponse(w, http.StatusCreated, json)
}
