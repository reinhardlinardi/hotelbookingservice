package main

import (
	"log"
	"net/http"

	hotelbooking "./hotelbooking"
	"github.com/julienschmidt/httprouter"
)

func main() {
	hotelbooking.ConnectDB()
	defer hotelbooking.DisconnectDB()

	hotelbooking.GetHotelAuthToken()
	router := httprouter.New()

	router.POST("/customer", hotelbooking.RegisterProfile)

	log.Fatal(http.ListenAndServe(":8060", router))
}
