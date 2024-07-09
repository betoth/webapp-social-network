package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp-social-network/src/config"
	"webapp-social-network/src/cookie"
	"webapp-social-network/src/models"
	"webapp-social-network/src/requests"
	"webapp-social-network/src/response"
	"webapp-social-network/src/utils"

	"github.com/gorilla/mux"
)

// LoadingLoginPage loading a login page
func LoadingLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookie.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

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

	cookie, err := cookie.Read(r)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: err.Error()})
		return
	}

	userID, err := strconv.ParseUint(cookie["id"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "home.html", struct {
		Publications []models.Publication
		UserID       uint64
	}{
		Publications: publications,
		UserID:       userID,
	})
}

// LoadUpdatePublicationPage loading update publication page
func LoadUpdatePublicationPage(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameter["publicationId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/%d", config.APIURL, publicationID)
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

	var publication models.Publication
	err = json.NewDecoder(resp.Body).Decode(&publication)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "update-publication.html", publication)
}

// LoadUsersPage loading user page
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrNick)
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

	var users []models.User

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: err.Error()})
		return
	}
	utils.ExecuteTemplate(w, "users.html", users)

}

// LoadUsersProfilePage load a user profile page
func LoadUsersProfilePage(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)
	userID, err := strconv.ParseUint(parameter["userId"], 10, 64)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookie.Read(r)
	loggedUser, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == loggedUser {
		http.Redirect(w, r, "/profile", 302)
		return
	}

	user, err := models.SearchFullUser(userID, r)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		LoggedUserID uint64
	}{
		User:         user,
		LoggedUserID: loggedUser,
	})
}

// LoadUsersLoggedProfilePage loading a serl profile page
func LoadUsersLoggedProfilePage(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookie.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.SearchFullUser(userID, r)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "profile.html", user)

}

// LoadUsersUpdatePage loading a update user page
func LoadUsersUpdatePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookie.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.SearchUserData(userID, r)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "update-user.html", user)
}

// LoadUpdatePasswordPage load update password page
func LoadUpdatePasswordPage(w http.ResponseWriter, r *http.Request) {

	utils.ExecuteTemplate(w, "update-password.html", nil)
}
