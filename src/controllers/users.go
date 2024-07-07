package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp-social-network/src/config"
	"webapp-social-network/src/response"
)

// CreateUser create a user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":      r.FormValue("name"),
		"nick_name": r.FormValue("nick_name"),
		"email":     r.FormValue("email"),
		"password":  r.FormValue("password"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.TreatStatusCodeError(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, nil)
}
