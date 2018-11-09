package hotelbooking

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

/* Types */

type Agent struct {
	Name       *string `json:"name"`
	Address    *string `json:"address"`
	ExpireDate *string `json:"expire_date"`
}

type RegisterAgentResponseData struct {
	ID int `json:"id"`
}

/* API */

func RegisterAgent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var agent Agent
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("RegisterAgent :", err)

		SendBadRequestWithData(w)
		return
	}

	err = json.Unmarshal(body, &agent)

	if err != nil {
		log.Println("RegisterAgent :", err)

		SendBadRequestWithData(w)
		return
	}

	if agent.Name == nil || agent.Address == nil || agent.ExpireDate == nil {
		SendBadRequestWithData(w)
		return
	}

	expireDate, err := time.Parse("02-01-2006", *agent.ExpireDate)

	if err != nil {
		SendBadRequestWithData(w)
		return
	}

	statement, err := db.Prepare("INSERT INTO agent VALUES (0, ?, ?, ?)")

	if err != nil {
		log.Println("RegisterAgent :", err)
		return
	}

	defer statement.Close()

	res, err := statement.Exec(*agent.Name, *agent.Address, expireDate)

	if err != nil {
		log.Println("RegisterAgent :", err)
		return
	}

	id, _ := res.LastInsertId()

	data := RegisterAgentResponseData{
		ID: int(id),
	}

	SendOKWithData(w, data)
}
