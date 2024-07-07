package requests

import (
	"io"
	"net/http"
	"webapp-social-network/src/cookie"
)

// MakeRequestWithAuthentication add token to requisition
func MakeRequestWithAuthentication(r *http.Request, methodo, url string, data io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(methodo, url, data)
	if err != nil {
		return nil, err
	}

	cookie, err := cookie.Read(r)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
