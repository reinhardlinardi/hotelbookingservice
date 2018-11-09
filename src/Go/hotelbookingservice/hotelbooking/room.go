package hotelbooking

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

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

/* Helper */

func isRoomTypeValid(roomType string) bool {
	roomType = strings.ToLower(roomType)

	if roomType == "single" || roomType == "double" || roomType == "family" {
		return true
	} else {
		return false
	}
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

func GetRoomInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("GetRoomInfo :", err)

		SendNotFoundWithData(w)
		return
	}

	statement, err := db.Prepare("SELECT type, description, tv, ac, internet, water, refrigerator, deposit_box, wardrobe, window, balcony, price FROM room WHERE id = ?")

	if err != nil {
		log.Println("GetRoomInfo :", err)

		SendNotFoundWithData(w)
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

		SendNotFound(w)
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
