package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aixoio/bridge/server/db"
)

type signup_request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var dat signup_request
	err := json.NewDecoder(r.Body).Decode(&dat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(dat.Username) > 32 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.SignUpUser(dat.Username, dat.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
