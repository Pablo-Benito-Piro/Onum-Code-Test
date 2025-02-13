package commons

import (
	"net/http"
)

func SendResponse(writer http.ResponseWriter, status int, data []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_, _ = writer.Write(data)
}
func SendError(data []byte, writer http.ResponseWriter, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_, _ = writer.Write(data)
}

func SendAuctionEndDateError(writer http.ResponseWriter, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_, _ = writer.Write([]byte("Cannot create an auction with an end date before that now"))
}
func SendBidEndDateError(writer http.ResponseWriter, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_, _ = writer.Write([]byte("Cannot create an Bid when an auction is close "))
}
func SendADuplicateAuctionError(data string, writer http.ResponseWriter, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_, _ = writer.Write([]byte("There is already an " + data + "with this id"))
}
