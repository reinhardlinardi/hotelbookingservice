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
	ID    int    `json:"id"`
	Token string `json:"token"`
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

		SendBadRequest(w, true)
		return
	}

	err = json.Unmarshal(body, &profile)

	if err != nil {
		log.Println("RegisterProfile :", err)

		SendBadRequest(w, true)
		return
	}

	if profile.Name == "" || profile.ID == "" || profile.Email == "" {
		SendBadRequest(w, true)
		return
	}

	token := GenerateToken(profile.Email)
	authStatement, err := db.Prepare("INSERT INTO auth VALUES (0, ?, 'customer')")

	if err != nil {
		log.Println("RegisterProfile :", err)
		return
	}

	defer authStatement.Close()

	_, err = authStatement.Exec(token)

	if err != nil {
		log.Println("RegisterProfile :", err)
		return
	}

	customerStatement, err := db.Prepare("INSERT INTO customer VALUES (0, ?, ?, ?, ?)")

	if err != nil {
		log.Println("RegisterProfile :", err)
		return
	}

	defer customerStatement.Close()

	res, err := customerStatement.Exec(profile.Name, profile.ID, profile.Email, token)

	if err != nil {
		log.Println("RegisterProfile :", err)
		return
	}

	id, _ := res.LastInsertId()

	data := RegisterProfileResponseData{
		ID:    int(id),
		Token: token,
	}

	SendOK(w, data)
}
