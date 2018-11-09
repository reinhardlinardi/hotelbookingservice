package hotelbooking

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

func UpdateEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var employee Employee
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("UpdateEmployee :", err)

		SendBadRequest(w)
		return
	}

	err = json.Unmarshal(body, &employee)

	if err != nil {
		log.Println("UpdateEmployee :", err)

		SendBadRequest(w)
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("UpdateEmployee :", err)

		SendNotFound(w)
		return
	}

	statement, err := db.Prepare("SELECT name, address, job_desc FROM employee WHERE id = ?")

	if err != nil {
		log.Println("UpdateEmployee :", err)

		SendNotFound(w)
		return
	}

	defer statement.Close()

	var employeeData Employee
	err = statement.QueryRow(id).Scan(&employeeData.Name, &employeeData.Address, &employeeData.JobDesc)

	if err != nil {
		log.Println("UpdateEmployee :", err)

		SendNotFound(w)
		return
	}

	if employee.Name == nil {
		employee.Name = employeeData.Name
	}

	if employee.Address == nil {
		employee.Address = employeeData.Address
	}

	if employee.JobDesc == nil {
		employee.JobDesc = employeeData.JobDesc
	}

	statement, err = db.Prepare("UPDATE employee SET name = ?, address = ?, job_desc = ? WHERE id = ?")

	if err != nil {
		log.Println("UpdateEmployee :", err)
		return
	}

	_, err = statement.Exec(*employee.Name, *employee.Address, *employee.JobDesc, id)

	if err != nil {
		log.Println("UpdateEmployee :", err)
		return
	}

	SendOK(w)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		log.Println("DeleteEmployee :", err)

		SendNotFound(w)
		return
	}

	statement, err := db.Prepare("SELECT id FROM employee WHERE id = ?")

	if err != nil {
		log.Println("DeleteEmployee :", err)

		SendNotFound(w)
		return
	}

	defer statement.Close()

	var dummy int
	err = statement.QueryRow(id).Scan(&dummy)

	if err != nil {
		log.Println("DeleteEmployee :", err)

		SendNotFound(w)
		return
	}

	statement, err = db.Prepare("DELETE FROM employee WHERE id = ?")

	if err != nil {
		log.Println("DeleteEmployee :", err)
		return
	}

	_, err = statement.Exec(id)

	if err != nil {
		log.Println("DeleteEmployee :", err)
		return
	}

	SendOK(w)
}
