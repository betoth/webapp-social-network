package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp-social-network/src/config"
	"webapp-social-network/src/requests"
	"webapp-social-network/src/response"

	"github.com/gorilla/mux"
)

// CreatePublication create a publication
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}
	url := fmt.Sprintf("%s/publication", config.APIURL)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(publication))
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

// Like add one like in a publication
func Like(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameter["publicationId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/%d/like", config.APIURL, publicationID)
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

// Unlike add one like in a publication
func Unlike(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameter["publicationId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/%d/unlike", config.APIURL, publicationID)
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

// UpdatePublication create a publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameter["publicationId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	r.ParseForm()
	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/%d", config.APIURL, publicationID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodPut, url, bytes.NewBuffer(publication))
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

// DeletePublication delete a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameter["publicationId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/%d", config.APIURL, publicationID)
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

	response.JSON(w, resp.StatusCode, nil)
}
