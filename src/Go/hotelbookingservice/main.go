package main

import (
	"log"
	"net/http"

	hotelbooking "./hotelbooking"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Customer - Public
	router.POST("/customer", hotelbooking.RegisterProfile)

	log.Fatal(http.ListenAndServe(":8060", router))
}
