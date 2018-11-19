package hotel

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/* Variables */
var (
	db *sql.DB
)

/* Types */

type Config struct {
	MySQL struct {
		User string `json:"user"`
		Pass string `json:"pass"`
		DB   string `json:"db"`
	} `json:"mysql"`
}

/* Functions */

func Init() {
	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Fatal("ConnectDB : ", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)

	if err != nil {
		log.Fatal("ConnectDB : ", err)
	}

	connStr := fmt.Sprintf("%s:%s@/%s", config.MySQL.User, config.MySQL.Pass, config.MySQL.DB)
	conn, err := sql.Open("mysql", connStr)

	if err != nil {
		log.Fatal("ConnectDB : ", err)
	}

	db = conn
}

func CleanUp() {
	db.Close()
}
