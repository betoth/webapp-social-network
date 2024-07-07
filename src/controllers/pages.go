package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp-social-network/src/config"
	"webapp-social-network/src/models"
	"webapp-social-network/src/requests"
	"webapp-social-network/src/response"
	"webapp-social-network/src/utils"
)

// LoadingLoginPage loading a login page
func LoadingLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadingUserRegisterPage loading a users register page
func LoadingUserRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)

}

// LoadingHomePage loading a home page
func LoadingHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.APIURL)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.TreatStatusCodeError(w, resp)
		return
	}

	var publications []models.Publication

	err = json.NewDecoder(resp.Body).Decode(&publications)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "home.html", publications)
}
