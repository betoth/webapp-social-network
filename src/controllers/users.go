package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp-social-network/src/config"
	"webapp-social-network/src/cookie"
	"webapp-social-network/src/requests"
	"webapp-social-network/src/response"

	"github.com/gorilla/mux"
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

// Unfollow unfollow a user
func Unfollow(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	userID, err := strconv.ParseUint(parameter["userId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodPost, url, nil)
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

// Follow unfollow a user
func Follow(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	userID, err := strconv.ParseUint(parameter["userId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodPost, url, nil)
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

// UpdateUser create a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"name":      r.FormValue("name"),
		"nick_name": r.FormValue("nick_name"),
		"email":     r.FormValue("email"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookie.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodPut, url, bytes.NewBuffer(user))
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

// UpdatePassword update user password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	passwords, err := json.Marshal(map[string]string{
		"current-password": r.FormValue("current"),
		"new-password":     r.FormValue("new"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}
	cookie, _ := cookie.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/update-password", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(passwords))
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

// DeleteUser delete a user perfil
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookie.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodDelete, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.TreatStatusCodeError(w, resp)
		return
	}

}
