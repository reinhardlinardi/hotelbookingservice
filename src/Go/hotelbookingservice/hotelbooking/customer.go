package hotelbooking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var profile Profile
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("RegisterProfile", err)
		return
	}

	err = json.Unmarshal(body, &profile)

	if err != nil {
		log.Println("RegisterProfile", err)
		return
	}

	fmt.Fprintln(w, profile.Name)
	fmt.Fprintln(w, profile.ID)
	fmt.Fprintln(w, profile.Email)
}
