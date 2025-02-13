package service

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"onumTest/commons"
	"onumTest/dao"
	"onumTest/models"
	"time"
)

func CreateBid(w http.ResponseWriter, r *http.Request) {
	bid := models.Bid{}
	err := json.NewDecoder(r.Body).Decode(&bid)
	if err != nil {
		commons.SendError([]byte(err.Error()), w, http.StatusBadRequest)
		return
	}

	auction := dao.FindAuctionByID(bid.AuctionID)

	if auction.ID != bid.AuctionID {

		commons.SendADuplicateAuctionError("Bid", w, http.StatusInternalServerError)
		return
	}

	if bid.Bid < auction.StartValue {
		commons.SendError([]byte("cannot create a bid lower than the Auction start value"), w, http.StatusInternalServerError)
		return
	}

	if auction.BidEndTime < time.Now().UnixMilli() || auction.BidStartTime > time.Now().UnixMilli() {
		commons.SendBidEndDateError(w, http.StatusInternalServerError)
		return
	}
	bid.Status = models.Processing

	bid, err = dao.CreateBidDAO(bid)

	statusProcess(bid)

	if err != nil {
		commons.SendError([]byte(err.Error()), w, http.StatusInternalServerError)
		return
	}

	commons.SendResponse(w, http.StatusCreated, []byte("se ha procesado la puja"))
}

func notifyClient(bid models.Bid) {

	jsonMarshal, err := json.Marshal(bid)

	if err != nil {
		log.Println("Error creating the notify", err)
		return
	}
	req, err := http.NewRequest(http.MethodPut, bid.Update, bytes.NewReader(jsonMarshal))

	if err != nil {
		log.Println("Error creating the notify", err)
		return
	}

	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending the updated Bid:", err)
		return
	}
	defer resp.Body.Close()
}

func statusProcess(lastBid models.Bid) {
	bids := dao.FindBids(lastBid.AuctionID)
	if len(bids) == 1 {
		lastBid.Status = models.Best
		dao.SaveBid(lastBid)
		notifyClient(lastBid)

	} else {
		bestBid := dao.FindBidsByStatusBest()
		if bestBid.Bid >= lastBid.Bid {
			lastBid.Status = models.Outbided
			dao.SaveBid(lastBid)
			notifyClient(lastBid)

		} else {
			lastBid.Status = models.Best
			bestBid.Status = models.Outbided
			dao.SaveBid(lastBid)
			notifyClient(lastBid)
			dao.SaveBid(bestBid)
			notifyClient(bestBid)
		}

	}

}

func ClientSimulatorCallBack(w http.ResponseWriter, r *http.Request) {
	bid := models.Bid{}
	err := json.NewDecoder(r.Body).Decode(&bid)
	if err != nil {
		commons.SendError([]byte(err.Error()), w, http.StatusBadRequest)
	}
	result, err := json.Marshal(bid)
	if err != nil {
		commons.SendError([]byte(err.Error()), w, http.StatusBadRequest)
	}
	commons.SendResponse(w, http.StatusCreated, result)
}

func GetBidsByAuctionIDandyClientID(w http.ResponseWriter, r *http.Request) {
	var bids []models.Bid
	bidSearch := models.BidSearch{}
	body, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal(body, &bidSearch); err != nil {
		commons.SendError([]byte(err.Error()), w, http.StatusNotFound)
		return
	}

	db := commons.GetConnection()
	defer db.Close()

	bids = dao.FindBidsByAuctionIDAndClientID(bidSearch.ClientID, bidSearch.AuctionID)

	if len(bids) == 0 {
		commons.SendError([]byte("there is no bid with that client id and auction id"), w, http.StatusNotFound)
	} else {
		result, err := json.Marshal(bids)
		if err != nil {
			commons.SendError([]byte(err.Error()), w, http.StatusInternalServerError)
		}
		commons.SendResponse(w, http.StatusOK, result)
	}

}
