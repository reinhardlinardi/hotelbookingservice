package hotel

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/* Types */

type Profile struct {
	Name  *string `json:"name"`
	ID    *string `json:"id"`
	Email *string `json:"email"`
}

type RegisterProfileResponseData struct {
	ID int `json:"id"`
}

type GetIDByProfileResponseData struct {
	ID int `json:"id"`
}

type GetProfileByIDResponseData struct {
	Identity int `json:"Identity"`
	Name int `json:"Name"`
	Email int `json:"Email"`
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

	if profile.Name == nil || profile.ID == nil || profile.Email == nil {
		SendBadRequestWithData(w)
		return
	}

	statement, err := db.Prepare("INSERT INTO customer VALUES (0, ?, ?, ?)")

	if err != nil {
		log.Println("RegisterProfile :", err)
		return
	}

	defer statement.Close()

	res, err := statement.Exec(*profile.Name, *profile.ID, *profile.Email)

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

func GetIDByProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.URL.Query().Get("name")
	identity := r.URL.Query().Get("id")
	email := r.URL.Query().Get("email")

	if name == "" || identity == "" || email == "" {
		SendBadRequestWithData(w)
		return
	}

	statement, err := db.Prepare("SELECT id FROM customer WHERE name = ? AND identity = ? AND email = ?")

	if err != nil {
		log.Println("GetIDByProfile :", err)
		return
	}

	defer statement.Close()

	var data GetIDByProfileResponseData
	err = statement.QueryRow(name, identity, email).Scan(&data.ID)

	if err != nil {
		log.Println("GetIDByProfile :", err)

		SendOKWithData(w, nil)
		return
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

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("UpdateProfile :", err)

		SendNotFound(w)
		return
	}

	statement, err := db.Prepare("SELECT name, identity, email FROM customer WHERE id = ?")

	if err != nil {
		log.Println("UpdateProfile :", err)
		return
	}

	defer statement.Close()

	var profileData Profile
	err = statement.QueryRow(id).Scan(&profileData.Name, &profileData.ID, &profileData.Email)

	if err != nil {
		log.Println("UpdateProfile :", err)

		SendNotFound(w)
		return
	}

	if profile.Name == nil {
		profile.Name = profileData.Name
	}

	if profile.ID == nil {
		profile.ID = profileData.ID
	}

	if profile.Email == nil {
		profile.Email = profileData.Email
	}

	statement, err = db.Prepare("UPDATE customer SET name = ?, identity = ?, email = ? WHERE id = ?")

	if err != nil {
		log.Println("UpdateProfile :", err)
		return
	}

	_, err = statement.Exec(*profile.Name, *profile.ID, *profile.Email, id)

	if err != nil {
		log.Println("UpdateProfile :", err)
		return
	}

	SendOK(w)
}

//NEW CISCO
func GetProfileByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("UpdateProfile :", err)

		SendNotFound(w)
		return
	}

	statement, err := db.Prepare("SELECT name,identity,email FROM customer WHERE id = ?")

	if err != nil {
		log.Println("GetIDByProfile :", err)
		return
	}

	defer statement.Close()

	var profileData GetProfileByIDResponseData
	err = statement.QueryRow(id).Scan(&profileData.Name, &profileData.Identity, &profileData.Email)

	if err != nil {
		log.Println("GetProfileByID :", err)

		SendOKWithData(w, nil)
		return
	}

	SendOKWithData(w, profileData)
}