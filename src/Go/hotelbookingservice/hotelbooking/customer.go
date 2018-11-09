package hotelbooking

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* Types */

type Profile struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Email string `json:"email"`
}

type RegisterProfileResponseData struct {
	ID int `json:"id"`
}

type RegisterProfileResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Data    RegisterProfileResponseData `json:"data"`
}

/* API */

func RegisterProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var profile Profile
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("RegisterProfile :", err)

		SendBadRequestWithData(w)
		return
	}

	err = json.Unmarshal(body, &profile)

	if err != nil {
		log.Println("RegisterProfile :", err)

		SendBadRequestWithData(w)
		return
	}

	if profile.Name == "" || profile.ID == "" || profile.Email == "" {
		SendBadRequestWithData(w)
		return
	}

	statement, err := db.Prepare("INSERT INTO customer VALUES (0, ?, ?, ?)")

	if err != nil {
		log.Println("RegisterProfile :", err)
		return
	}

	defer statement.Close()

	res, err := statement.Exec(profile.Name, profile.ID, profile.Email)

	if err != nil {
		log.Println("RegisterProfile :", err)
		return
	}

	id, _ := res.LastInsertId()

	data := RegisterProfileResponseData{
		ID: int(id),
	}

	SendOKWithData(w, data)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var profile Profile
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("UpdateProfile :", err)

		SendBadRequest(w)
		return
	}

	err = json.Unmarshal(body, &profile)

	if err != nil {
		log.Println("UpdateProfile :", err)

		SendBadRequest(w)
		return
	}
	/*
		id, err := strconv.Atoi(ps.ByName("id"))

		if err != nil {
			log.Println("UpdateProfile :", err)

			SendNotFound(w)
			return
		}
	*/
	if profile.Name == "" && profile.ID == "" && profile.Email == "" {
		SendOK(w)
		return
	}
}
