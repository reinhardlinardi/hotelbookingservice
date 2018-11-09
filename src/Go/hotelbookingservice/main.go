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

	router := httprouter.New()

	router.POST("/customer", hotelbooking.RegisterProfile)
	router.PUT("/customer/:id", hotelbooking.UpdateProfile)

	router.POST("/employee", hotelbooking.AddEmployee)
	router.PUT("/employee/:id", hotelbooking.UpdateEmployee)
	router.DELETE("/employee/:id", hotelbooking.DeleteEmployee)

	log.Fatal(http.ListenAndServe(":8060", router))
}
