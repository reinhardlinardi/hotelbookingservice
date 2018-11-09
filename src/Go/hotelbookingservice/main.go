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

	router.POST("/agent", hotelbooking.RegisterAgent)
	router.PUT("/agent/:id", hotelbooking.UpdateAgent)
	router.DELETE("/agent/:id", hotelbooking.DeleteAgent)

	router.POST("/room", hotelbooking.CreateRoom)
	router.GET("/room/:id", hotelbooking.GetRoomInfo)
	router.DELETE("/room/:id", hotelbooking.DeleteRoom)

	router.POST("/invoice", hotelbooking.CreateInvoice)
	router.PUT("/invoice/:id", hotelbooking.UpdateInvoicePaymentStatus)

	log.Fatal(http.ListenAndServe(":8060", router))
}
