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
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

/* Functions */

func SendOK(w http.ResponseWriter, data interface{}) {
	response := Response{
		Success: true,
		Message: "OK",
		Data:    data,
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

func SendBadRequest(w http.ResponseWriter, paramRequired bool) {
	var message string

	if paramRequired {
		message = ERR_BAD_REQUEST_REQUIRED
	} else {
		message = ERR_BAD_REQUEST_OPTIONAL
	}

	response := Response{
		Success: false,
		Message: message,
		Data:    nil,
	}

	json, err := json.Marshal(response)

	if err != nil {
		log.Println("SendBadRequest :", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(json)
}
