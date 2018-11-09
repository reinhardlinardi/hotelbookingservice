package hotelbooking

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

func UpdateAgent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var agent Agent
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("UpdateAgent :", err)

		SendBadRequest(w)
		return
	}

	err = json.Unmarshal(body, &agent)

	if err != nil {
		log.Println("UpdateAgent :", err)

		SendBadRequest(w)
		return
	}

	expireDate, err := time.Parse("02-01-2006", *agent.ExpireDate)

	if err != nil {
		log.Println("UpdateAgent :", err)

		SendBadRequest(w)
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("UpdateAgent :", err)

		SendNotFound(w)
		return
	}

	statement, err := db.Prepare("SELECT name, address, expire_date FROM agent WHERE id = ?")

	if err != nil {
		log.Println("UpdateAgent :", err)
		return
	}

	defer statement.Close()

	var agentData Agent
	err = statement.QueryRow(id).Scan(&agentData.Name, &agentData.Address, &agentData.ExpireDate)

	if err != nil {
		log.Println("UpdateAgent :", err)

		SendNotFound(w)
		return
	}

	if agent.Name == nil {
		agent.Name = agentData.Name
	}

	if agent.Address == nil {
		agent.Address = agentData.Address
	}

	if agent.ExpireDate == nil {
		agent.ExpireDate = agentData.ExpireDate
	}

	statement, err = db.Prepare("UPDATE agent SET name = ?, address = ?, expire_date = ? WHERE id = ?")

	if err != nil {
		log.Println("UpdateAgent :", err)
		return
	}

	_, err = statement.Exec(*agent.Name, *agent.Address, expireDate, id)

	if err != nil {
		log.Println("UpdateAgent :", err)
		return
	}

	SendOK(w)
}

func DeleteAgent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("DeleteAgent :", err)

		SendNotFound(w)
		return
	}

	statement, err := db.Prepare("SELECT id FROM agent WHERE id = ?")

	if err != nil {
		log.Println("DeleteAgent :", err)

		SendNotFound(w)
		return
	}

	defer statement.Close()

	var dummy int
	err = statement.QueryRow(id).Scan(&dummy)

	if err != nil {
		log.Println("DeleteAgent :", err)

		SendNotFound(w)
		return
	}

	statement, err = db.Prepare("DELETE FROM agent WHERE id = ?")

	if err != nil {
		log.Println("DeleteAgent :", err)
		return
	}

	_, err = statement.Exec(id)

	if err != nil {
		log.Println("DeleteAgent :", err)
		return
	}

	SendOK(w)
}
