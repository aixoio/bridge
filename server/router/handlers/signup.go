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

type signup_responce struct {
	Success bool   `json:"success"`
	Err_msg string `json:"errMsg"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var dat signup_request
	err := json.NewDecoder(r.Body).Decode(&dat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json_dat, _ := json.Marshal(signup_responce{
			Success: false,
			Err_msg: "Server side error, if your the admin please check the logs",
		})
		w.Write(json_dat)
		return
	}

	if len(dat.Username) > 32 {
		w.WriteHeader(http.StatusBadRequest)
		json_dat, _ := json.Marshal(signup_responce{
			Success: false,
			Err_msg: "Username is too long",
		})
		w.Write(json_dat)
		return
	}

	err = db.SignUpUser(dat.Username, dat.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json_dat, _ := json.Marshal(signup_responce{
			Success: false,
			Err_msg: "Server side error, if your the admin please check the logs",
		})
		w.Write(json_dat)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json_dat, _ := json.Marshal(signup_responce{
		Success: true,
		Err_msg: "",
	})
	w.Write(json_dat)

}
