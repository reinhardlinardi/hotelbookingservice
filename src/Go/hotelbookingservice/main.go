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
	router.GET("/customer", hotelbooking.GetIDByProfile)
	router.PUT("/customer/:id", hotelbooking.UpdateProfile)

	router.POST("/employee", hotelbooking.AddEmployee)
	router.PUT("/employee/:id", hotelbooking.UpdateEmployee)
	router.DELETE("/employee/:id", hotelbooking.DeleteEmployee)

	router.POST("/agent", hotelbooking.RegisterAgent)
	router.PUT("/agent/:id", hotelbooking.UpdateAgent)
	router.DELETE("/agent/:id", hotelbooking.DeleteAgent)

	router.POST("/room", hotelbooking.CreateRoom)
	router.GET("/room", hotelbooking.GetAvailableRoomsData)
	router.GET("/room/:id", hotelbooking.GetRoomData)
	router.DELETE("/room/:id", hotelbooking.DeleteRoom)

	router.POST("/invoice", hotelbooking.CreateInvoice)
	router.GET("/invoice/:id", hotelbooking.GetInvoice)
	router.PUT("/invoice/:id", hotelbooking.UpdateInvoicePaymentStatus)
	router.DELETE("/invoice/:id", hotelbooking.CancelInvoice)

	log.Fatal(http.ListenAndServe(":8060", router))
}
