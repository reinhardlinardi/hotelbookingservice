package hotelbooking

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* Types */

type Employee struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
	JobDesc *string `json:"job_desc"`
}

type AddEmployeeResponseData struct {
	ID int `json:"id"`
}

/* API */

func AddEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var employee Employee
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("AddEmployee :", err)

		SendBadRequestWithData(w)
		return
	}

	err = json.Unmarshal(body, &employee)

	if err != nil {
		log.Println("AddEmployee :", err)

		SendBadRequestWithData(w)
		return
	}

	if employee.Name == nil || employee.Address == nil || employee.JobDesc == nil {
		SendBadRequestWithData(w)
		return
	}

	statement, err := db.Prepare("INSERT INTO employee VALUES (0, ?, ?, ?)")

	if err != nil {
		log.Println("AddEmployee :", err)
		return
	}

	defer statement.Close()

	res, err := statement.Exec(*employee.Name, *employee.Address, *employee.JobDesc)

	if err != nil {
		log.Println("AddEmployee :", err)
		return
	}

	id, _ := res.LastInsertId()

	data := AddEmployeeResponseData{
		ID: int(id),
	}

	SendOKWithData(w, data)
}
