package hotelbooking

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/* Variables */
var (
	db *sql.DB
)

/* Types */

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

/* Functions */

func ConnectDB() {
	data, err := ioutil.ReadFile("config/dbconfig.json")

	if err != nil {
		log.Fatal("ConnectDB : ", err)
	}

	var config DBConfig
	err = json.Unmarshal(data, &config)

	if err != nil {
		log.Fatal("ConnectDB : ", err)
	}

	connStr := config.Username + ":" + config.Password + "@/" + config.Database
	conn, err := sql.Open("mysql", connStr)

	if err != nil {
		log.Fatal("ConnectDB : ", err)
	}

	db = conn
}

func DisconnectDB() {
	db.Close()
}
