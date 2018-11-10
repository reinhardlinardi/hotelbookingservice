package hotelbooking

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

/* Types */

type Room struct {
	RoomType       *string `json:"room_type"`
	Description    *string `json:"description"`
	TV             *bool   `json:"tv"`
	AC             *bool   `json:"ac"`
	Internet       *bool   `json:"internet"`
	HotWater       *bool   `json:"hot_water"`
	Refrigerator   *bool   `json:"refrigerator"`
	SafeDepositBox *bool   `json:"safe_deposit_box"`
	Wardrobe       *bool   `json:"wardrobe"`
	Window         *bool   `json:"window"`
	Balcony        *bool   `json:"balcony"`
	Price          *int    `json:"price"`
}

type CreateRoomResponseData struct {
	ID int `json:"ID"`
}

type GetRoomAvailabilityResponseData struct {
	Available bool `json:"available"`
}

/* Helper */

func isRoomTypeValid(roomType string) bool {
	roomType = strings.ToLower(roomType)

	if roomType == "single" || roomType == "double" || roomType == "family" {
		return true
	}

	return false
}

/* API */

func CreateRoom(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var room Room
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("CreateRoom :", err)

		SendBadRequestWithData(w)
		return
	}

	err = json.Unmarshal(body, &room)

	if err != nil {
		log.Println("CreateRoom :", err)

		SendBadRequestWithData(w)
		return
	}

	if room.RoomType == nil || room.Description == nil || room.TV == nil || room.AC == nil || room.Internet == nil || room.HotWater == nil || room.Refrigerator == nil || room.SafeDepositBox == nil || room.Wardrobe == nil || room.Window == nil || room.Balcony == nil || room.Price == nil {
		SendBadRequestWithData(w)
		return
	}

	if !isRoomTypeValid(*room.RoomType) {
		SendBadRequestWithData(w)
		return
	}

	statement, err := db.Prepare("INSERT INTO room VALUES (0, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Println("CreateRoom :", err)
		return
	}

	defer statement.Close()

	res, err := statement.Exec(*room.RoomType, *room.Description, *room.TV, *room.AC, *room.Internet, *room.HotWater, *room.Refrigerator, *room.SafeDepositBox, *room.Wardrobe, *room.Window, *room.Balcony, *room.Price)

	if err != nil {
		log.Println("CreateRoom :", err)
		return
	}

	id, _ := res.LastInsertId()

	data := CreateRoomResponseData{
		ID: int(id),
	}

	SendOKWithData(w, data)
}

/*
func GetAvailableRooms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	checkInStr := r.URL.Query().Get("in")
	checkOutStr := r.URL.Query().Get("out")

	checkIn, err := time.Parse("02-01-2006", checkInStr)

	if err != nil {
		SendBadRequestWithData(w)
		return
	}

	checkOut, err := time.Parse("02-01-2006", checkOutStr)

	if err != nil {
		SendBadRequestWithData(w)
		return
	}

	if !checkIn.Before(checkOut) {
		SendBadRequestWithData(w)
		return
	}

	statement, err := db.Prepare("SELECT id FROM room")

	if err != nil {
		log.Println("GetAvailableRooms :", err)
		return
	}

	// Buat masing2 room ID, panggil API buat cek room sekian available dari tanggal check in smpe tanggal check out
}
*/
func GetRoomInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	re, _ := regexp.Compile("\\?.+$")

	if re.MatchString(r.URL.String()) {
		GetRoomAvailability(w, r, ps)
	} else {
		id, err := strconv.Atoi(ps.ByName("id"))

		if err != nil {
			log.Println("GetRoomInfo :", err)

			SendNotFoundWithData(w)
			return
		}

		statement, err := db.Prepare("SELECT type, description, tv, ac, internet, water, refrigerator, deposit_box, wardrobe, window, balcony, price FROM room WHERE id = ?")

		if err != nil {
			log.Println("GetRoomInfo :", err)
			return
		}

		defer statement.Close()

		var room Room
		err = statement.QueryRow(id).Scan(&room.RoomType, &room.Description, &room.TV, &room.AC, &room.Internet, &room.HotWater, &room.Refrigerator, &room.SafeDepositBox, &room.Wardrobe, &room.Window, &room.Balcony, &room.Price)

		if err != nil {
			log.Println("GetRoomInfo :", err)

			SendNotFoundWithData(w)
			return
		}

		SendOKWithData(w, room)
	}
}

func DeleteRoom(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("DeleteRoom :", err)

		SendNotFound(w)
		return
	}

	statement, err := db.Prepare("SELECT id FROM room WHERE id = ?")

	if err != nil {
		log.Println("DeleteRoom :", err)
		return
	}

	defer statement.Close()

	var dummy int
	err = statement.QueryRow(id).Scan(&dummy)

	if err != nil {
		log.Println("DeleteRoom :", err)

		SendNotFound(w)
		return
	}

	statement, err = db.Prepare("DELETE FROM room WHERE id = ?")

	if err != nil {
		log.Println("DeleteRoom :", err)
		return
	}

	_, err = statement.Exec(id)

	if err != nil {
		log.Println("DeleteRoom :", err)
		return
	}

	SendOK(w)
}

func GetRoomAvailability(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	checkInStr := r.URL.Query().Get("in")
	checkOutStr := r.URL.Query().Get("out")

	checkIn, err := time.Parse("02-01-2006", checkInStr)

	if err != nil {
		SendBadRequestWithData(w)
		return
	}

	checkOut, err := time.Parse("02-01-2006", checkOutStr)

	if err != nil {
		SendBadRequestWithData(w)
		return
	}

	if !checkIn.Before(checkOut) {
		SendBadRequestWithData(w)
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("GetRoomAvailability :", err)

		SendNotFoundWithData(w)
		return
	}

	statement, err := db.Prepare("SELECT id FROM room WHERE id = ?")

	if err != nil {
		log.Println("GetRoomAvailability :", err)
		return
	}

	defer statement.Close()

	var dummy int
	err = statement.QueryRow(id).Scan(&dummy)

	if err != nil {
		log.Println("GetRoomAvailability :", err)

		SendNotFoundWithData(w)
		return
	}

	statement, err = db.Prepare("SELECT in_date, out_date FROM invoice WHERE room_id = ? AND cancelled = 0 ORDER BY in_date ASC")

	if err != nil {
		log.Println("GetRoomAvailability :", err)
		return
	}

	var invoiceInStr, invoiceOutStr string
	var invoiceIn, invoiceOut time.Time
	var available bool

	rows, _ := statement.Query(id)

	for rows.Next() {
		err = rows.Scan(&invoiceInStr, &invoiceOutStr)

		if err != nil {
			log.Println("GetRoomAvailability :", err)
			return
		}

		invoiceIn, _ = time.Parse("2006-01-02", invoiceInStr)
		invoiceOut, _ = time.Parse("2006-01-02", invoiceOutStr)

		if checkIn.Before(invoiceIn) && (checkOut.Before(invoiceIn) || checkOut.Equal(invoiceIn)) {
			available = true
			break
		} else if checkIn.Before(invoiceIn) && checkOut.After(invoiceIn) {
			available = false
			break
		} else if (checkIn.Equal(invoiceIn) || checkIn.After(invoiceIn)) && checkIn.Before(invoiceOut) {
			available = false
			break
		}
	}

	data := GetRoomAvailabilityResponseData{
		Available: available,
	}

	SendOKWithData(w, data)
}
