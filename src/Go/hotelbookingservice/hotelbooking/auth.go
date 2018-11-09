package hotelbooking

import (
	"fmt"
	"hash/crc32"
	"log"
	"strconv"
	"time"
)

/* Authentication constant */

const (
	AUTHENTICATION_HOTEL = 1 << iota
	AUTHENTICATION_AGENT
	AUTHENTICATION_CUSTOMER
)

/* Variables */

var (
	hotelAuthToken string
)

/* Functions */

func GetHotelAuthToken() {
	statement, err := db.Prepare("SELECT token FROM auth WHERE type = ?")

	if err != nil {
		log.Println("GetHotelAuthToken :", err)
	}

	defer statement.Close()

	var token string
	err = statement.QueryRow("hotel").Scan(&token)

	if err != nil {
		log.Println("GetHotelAuthToken :", err)
	}

	hotelAuthToken = token
}

func GenerateToken(name string) string {
	currentMillis := time.Now().UnixNano() / int64(time.Millisecond)

	hash := crc32.ChecksumIEEE([]byte(name + ":" + strconv.FormatInt(currentMillis, 10)))
	token := fmt.Sprintf("%x", hash)

	return token
}
