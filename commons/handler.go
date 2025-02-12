package commons

import (
	"net/http"
)

func SendResponse(writer http.ResponseWriter, status int, data []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(data)
}
func SendResponseA(writer http.ResponseWriter, status int, data []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(data)

}
func SendError(data []byte, writer http.ResponseWriter, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(data)
}
func SendAuctionEndDateError(writer http.ResponseWriter, status int) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write([]byte("error fechas"))
}
func SendADuplicateAuctionError(writer http.ResponseWriter, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write([]byte("error"))
}
