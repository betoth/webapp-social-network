package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp-social-network/src/config"
	"webapp-social-network/src/cookie"
	"webapp-social-network/src/models"
	"webapp-social-network/src/response"
)

// Login make a login
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}
	url := fmt.Sprintf("%s/login", config.APIURL)
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

	var authenticationData models.AuthenticationData
	err = json.NewDecoder(resp.Body).Decode(&authenticationData)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: err.Error()})
		return

	}

	err = cookie.Save(w, authenticationData.ID, authenticationData.Token)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: err.Error()})
		return
	}

	response.JSON(w, resp.StatusCode, nil)

}
