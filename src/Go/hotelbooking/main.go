package main

import (
	"log"
	"net/http"

	hotel "./hotel"
	"github.com/julienschmidt/httprouter"
)

func main() {
	hotel.Init()
	defer hotel.CleanUp()

	router := httprouter.New()

	router.POST("/customer", hotel.RegisterProfile)
	router.GET("/customer", hotel.GetIDByProfile)
	router.PUT("/customer/:id", hotel.UpdateProfile)
	//UpdateByCisco
	router.GET("/customer/:id", hotel.GetProfileByID)

	router.POST("/employee", hotel.AddEmployee)
	router.PUT("/employee/:id", hotel.UpdateEmployee)
	router.DELETE("/employee/:id", hotel.DeleteEmployee)

	router.POST("/agent", hotel.RegisterAgent)
	router.PUT("/agent/:id", hotel.UpdateAgent)
	router.DELETE("/agent/:id", hotel.DeleteAgent)

	router.POST("/room", hotel.CreateRoom)
	router.GET("/room", hotel.GetAvailableRoomsData)
	router.GET("/room/:id", hotel.GetRoomData)
	router.DELETE("/room/:id", hotel.DeleteRoom)

	router.POST("/invoice", hotel.CreateInvoice)
	router.GET("/invoice/:id", hotel.GetInvoice)
	router.PUT("/invoice/:id", hotel.UpdateInvoicePaymentStatus)
	router.DELETE("/invoice/:id", hotel.CancelInvoice)

	log.Fatal(http.ListenAndServe(":8060", router))
}
