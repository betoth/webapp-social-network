package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webapp-social-network/src/config"
	"webapp-social-network/src/requests"
)

// User represents a user and social structure
type User struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	NickName     string        `json:"nick_name"`
	CreatedAt    time.Time     `json:"created_at"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publication"`
}

// SearchFullUser Fill in all user fields and return
func SearchFullUser(userID uint64, r *http.Request) (User, error) {
	channelUser := make(chan User, 1)
	channelFollowers := make(chan []User, 1)
	channelFollowing := make(chan []User, 1)
	channelPublication := make(chan []Publication, 1)
	errorChan := make(chan error, 4)

	go func() {
		user, err := SearchUserData(userID, r)
		if err != nil {
			errorChan <- err
		} else {
			channelUser <- user
		}
	}()

	go func() {
		followers, err := SearchFollowers(userID, r)
		if err != nil {
			errorChan <- err
		} else {
			channelFollowers <- followers
		}
	}()

	go func() {
		following, err := SearchFollowing(userID, r)
		if err != nil {
			errorChan <- err
		} else {
			channelFollowing <- following
		}
	}()

	go func() {
		publications, err := SearchPublications(userID, r)
		if err != nil {
			errorChan <- err
		} else {
			channelPublication <- publications
		}
	}()

	var (
		user        User
		followers   []User
		following   []User
		publication []Publication
	)
	var err error

	for i := 0; i < 4; i++ {
		select {
		case e := <-errorChan:
			err = e
		case userFilled := <-channelUser:
			user = userFilled
		case followersFilled := <-channelFollowers:
			followers = followersFilled
		case followingFilled := <-channelFollowing:
			following = followingFilled
		case publicationsFilled := <-channelPublication:
			publication = publicationsFilled
		}
	}

	if err != nil {
		return User{}, err
	}

	user.Followers = followers
	user.Following = following
	user.Publications = publication

	return user, nil
}

// SearchUserData fill user data
func SearchUserData(userID uint64, r *http.Request) (User, error) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		return User{}, fmt.Errorf("error fetching user data: %v", err)
	}
	defer resp.Body.Close()

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return User{}, fmt.Errorf("error decoding user data: %v", err)
	}

	return user, nil
}

// SearchFollowers fill followers data
func SearchFollowers(userID uint64, r *http.Request) ([]User, error) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching followers data: %v", err)
	}
	defer resp.Body.Close()

	var followers []User
	err = json.NewDecoder(resp.Body).Decode(&followers)
	if err != nil {
		return nil, fmt.Errorf("error decoding followers data: %v", err)
	}

	return followers, nil
}

// SearchFollowing fill following data
func SearchFollowing(userID uint64, r *http.Request) ([]User, error) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching following data: %v", err)
	}
	defer resp.Body.Close()

	var following []User
	err = json.NewDecoder(resp.Body).Decode(&following)
	if err != nil {
		return nil, fmt.Errorf("error decoding following data: %v", err)
	}

	return following, nil
}

// SearchPublications fill publications data
func SearchPublications(userID uint64, r *http.Request) ([]Publication, error) {
	url := fmt.Sprintf("%s/users/%d/publications", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching publications data: %v", err)
	}
	defer resp.Body.Close()

	var publications []Publication
	err = json.NewDecoder(resp.Body).Decode(&publications)
	if err != nil {
		return nil, fmt.Errorf("error decoding publications data: %v", err)
	}

	return publications, nil
}
