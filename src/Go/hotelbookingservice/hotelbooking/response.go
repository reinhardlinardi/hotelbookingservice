package hotelbooking

import (
	"encoding/json"
	"log"
	"net/http"
)

/* Messages */

const (
	OK                       = "OK"
	ERR_BAD_REQUEST_REQUIRED = "Missing required parameters or invalid data"
	ERR_BAD_REQUEST_OPTIONAL = "Invalid data"
	ERR_AUTHENTICATION       = "Invalid authentication"
	ERR_NOT_FOUND            = "Invalid ID"
)

/* Types */

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type DataResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

/* Functions */

func SendOK(w http.ResponseWriter) {
	response := Response{
		Success: true,
		Message: OK,
	}

	json, err := json.Marshal(response)

	if err != nil {
		log.Println("SendOK :", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func SendOKWithData(w http.ResponseWriter, data interface{}) {
	response := DataResponse{
		Success: true,
		Message: OK,
		Data:    data,
	}

	json, err := json.Marshal(response)

	if err != nil {
		log.Println("SendOKWithData :", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func SendBadRequest(w http.ResponseWriter) {
	response := Response{
		Success: false,
		Message: ERR_BAD_REQUEST_OPTIONAL,
	}

	json, err := json.Marshal(response)

	if err != nil {
		log.Println("SendBadRequest :", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func SendBadRequestWithData(w http.ResponseWriter) {
	response := DataResponse{
		Success: false,
		Message: ERR_BAD_REQUEST_REQUIRED,
		Data:    nil,
	}

	json, err := json.Marshal(response)

	if err != nil {
		log.Println("SendBadRequestWithData :", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(json)
}

func SendNotFound(w http.ResponseWriter) {
	response := Response{
		Success: false,
		Message: ERR_NOT_FOUND,
	}

	json, err := json.Marshal(response)

	if err != nil {
		log.Println("SendNotFound :", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(json)
}

func SendNotFoundWithData(w http.ResponseWriter) {
	response := DataResponse{
		Success: false,
		Message: ERR_NOT_FOUND,
		Data:    nil,
	}

	json, err := json.Marshal(response)

	if err != nil {
		log.Println("SendNotFoundWithData :", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(json)
}
