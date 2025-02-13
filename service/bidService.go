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
	error := json.NewDecoder(r.Body).Decode(&bid)
	if error != nil {
		commons.SendError([]byte(error.Error()), w, http.StatusBadRequest)
		return
	}

	auction := dao.FindAuctionByID(bid.AuctionID)

	if auction.ID != bid.AuctionID {

		commons.SendADuplicateAuctionError(w, http.StatusInternalServerError)
		return
	}

	if bid.Bid < auction.StartValue {
		commons.SendADuplicateAuctionError(w, http.StatusInternalServerError)
		return
	}

	if auction.BidEndTime < time.Now().UnixMilli() || auction.BidStartTime > time.Now().UnixMilli() {
		commons.SendADuplicateAuctionError(w, http.StatusInternalServerError)
		return
	}
	bid.Status = models.Processing

	bid, error = dao.CreateBidDAO(bid)

	statusProcess(bid)

	if error != nil {
		commons.SendError([]byte(error.Error()), w, http.StatusInternalServerError)
		return
	}

	commons.SendResponse(w, http.StatusCreated, []byte("se ha procesado la puja"))
}

func notifyClient(bid models.Bid) {

	json, _ := json.Marshal(bid)

	req, err := http.NewRequest(http.MethodPut, bid.Update, bytes.NewReader(json))

	if err != nil {
		log.Println("Error creando solicitud:", err)
		return
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error enviando actualización al cliente:", err)
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
	json.NewDecoder(r.Body).Decode(&bid)
	result, _ := json.Marshal(bid)
	commons.SendResponseA(w, http.StatusCreated, result)
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
		commons.SendAuctionEndDateError(w, http.StatusNotFound)
	} else {
		result, _ := json.Marshal(bids)
		commons.SendResponse(w, http.StatusOK, result)
	}

}
