package hotel

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

/* Types */

type Invoice struct {
	RoomID     *int    `json:"room_id"`
	CustomerID *int    `json:"customer_id"`
	CheckIn    *string `json:"in"`
	CheckOut   *string `json:"out"`
}

type CompleteInvoice struct {
	RoomID     int    `json:"room_id"`
	CustomerID int    `json:"customer_id"`
	CheckIn    string `json:"in"`
	CheckOut   string `json:"out"`
	Price      int    `json:"price"`
	Paid       bool   `json:"paid"`
	Cancelled  bool   `json:"cancelled"`
}

type CreateInvoiceResponseData struct {
	ID int `json:"id"`
}

/* API */

func CreateInvoice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var invoice Invoice
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("CreateInvoice :", err)

		SendBadRequestWithData(w)
		return
	}

	err = json.Unmarshal(body, &invoice)

	if err != nil {
		log.Println("CreateInvoice :", err)

		SendBadRequestWithData(w)
		return
	}

	if invoice.RoomID == nil || invoice.CustomerID == nil || invoice.CheckIn == nil || invoice.CheckOut == nil {
		SendBadRequestWithData(w)
		return
	}

	checkIn, err := time.Parse("02-01-2006", *invoice.CheckIn)

	if err != nil {
		SendBadRequestWithData(w)
		return
	}

	checkOut, err := time.Parse("02-01-2006", *invoice.CheckOut)

	if err != nil {
		SendBadRequestWithData(w)
		return
	}

	if !checkIn.Before(checkOut) {
		SendBadRequestWithData(w)
		return
	}

	var price, totalPrice, dummy int
	statement, err := db.Prepare("SELECT price FROM room WHERE id = ?")

	if err != nil {
		log.Println("CreateInvoice :", err)
		return
	}

	defer statement.Close()

	err = statement.QueryRow(*invoice.RoomID).Scan(&price)

	if err != nil {
		log.Println("CreateInvoice :", err)

		SendBadRequestWithData(w)
		return
	}

	statement, err = db.Prepare("SELECT id FROM customer WHERE id = ?")

	if err != nil {
		log.Println("CreateInvoice :", err)
		return
	}

	err = statement.QueryRow(*invoice.CustomerID).Scan(&dummy)

	if err != nil {
		log.Println("CreateInvoice :", err)

		SendBadRequestWithData(w)
		return
	}

	statement, err = db.Prepare("INSERT INTO invoice VALUES (0, ?, ?, ?, ?, ?, 0, 0)")

	if err != nil {
		log.Println("CreateInvoice :", err)
		return
	}

	totalPrice = (int(checkOut.Sub(checkIn).Hours()) / 24) * price
	res, err := statement.Exec(*invoice.RoomID, *invoice.CustomerID, checkIn, checkOut, totalPrice)

	if err != nil {
		log.Println("CreateInvoice :", err)
		return
	}

	id, _ := res.LastInsertId()

	data := CreateInvoiceResponseData{
		ID: int(id),
	}

	SendOKWithData(w, data)
}

func GetInvoice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("GetInvoice :", err)

		SendNotFoundWithData(w)
		return
	}

	statement, err := db.Prepare("SELECT room_id, customer_id, in_date, out_date, price, paid, cancelled FROM invoice WHERE id = ?")

	if err != nil {
		log.Println("GetInvoice :", err)
		return
	}

	defer statement.Close()

	var invoice CompleteInvoice
	err = statement.QueryRow(id).Scan(&invoice.RoomID, &invoice.CustomerID, &invoice.CheckIn, &invoice.CheckOut, &invoice.Price, &invoice.Paid, &invoice.Cancelled)

	if err != nil {
		log.Println("GetInvoice :", err)
		return
	}

	checkIn, err := time.Parse("2006-01-02", invoice.CheckIn)

	if err != nil {
		log.Println("GetInvoice :", err)
		return
	}

	checkOut, err := time.Parse("2006-01-02", invoice.CheckOut)

	if err != nil {
		log.Println("GetInvoice :", err)
		return
	}

	invoice.CheckIn = checkIn.Format("02-01-2006")
	invoice.CheckOut = checkOut.Format("02-01-2006")

	if err != nil {
		log.Println("GetInvoice :", err)

		SendNotFoundWithData(w)
		return
	}

	SendOKWithData(w, invoice)
}

func UpdateInvoicePaymentStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("UpdateInvoicePaymentStatus :", err)

		SendNotFound(w)
		return
	}

	statement, err := db.Prepare("SELECT id FROM invoice WHERE id = ?")

	if err != nil {
		log.Println("UpdateInvoicePaymentStatus :", err)
		return
	}

	defer statement.Close()

	var dummy int
	err = statement.QueryRow(id).Scan(&dummy)

	if err != nil {
		log.Println("UpdateInvoicePaymentStatus :", err)

		SendNotFound(w)
		return
	}

	statement, err = db.Prepare("UPDATE invoice SET paid = 1 WHERE id = ?")

	if err != nil {
		log.Println("UpdateInvoicePaymentStatus :", err)
		return
	}

	_, err = statement.Exec(id)

	if err != nil {
		log.Println("UpdateInvoicePaymentStatus :", err)
		return
	}

	SendOK(w)
}

func CancelInvoice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("CancelInvocie :", err)

		SendNotFound(w)
		return
	}

	statement, err := db.Prepare("SELECT id FROM invoice WHERE id = ?")

	if err != nil {
		log.Println("CancelInvoice :", err)
		return
	}

	defer statement.Close()

	var dummy int
	err = statement.QueryRow(id).Scan(&dummy)

	if err != nil {
		log.Println("CancelInvoice :", err)

		SendNotFound(w)
		return
	}

	statement, err = db.Prepare("UPDATE invoice SET cancelled = 1 WHERE id = ?")

	if err != nil {
		log.Println("CancelInvoice :", err)
		return
	}

	_, err = statement.Exec(id)

	if err != nil {
		log.Println("CancelInvoice :", err)
		return
	}

	SendOK(w)
}
