package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type AuthClient struct {
	Client       *http.Client
	BaseUrl      string
	AccessToken  string
	RefreshToken string
	CsrfToken    string
}

func NewAuthClient(baseURL string) (*AuthClient, error) {

	client := &http.Client{}

	auth := &AuthClient{
		Client:  client,
		BaseUrl: baseURL,
	}

	if err := auth.login(); err != nil {
		return nil, err
	}

	return auth, nil
}

func (a *AuthClient) login() error {
	loginURL := a.BaseUrl + "/auth/login"

	payload := map[string]string{
		"username": os.Getenv("ADMIN_USERNAME"),
		"password": os.Getenv("ADMIN_PASSWORD"),
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	for _, cookie := range resp.Cookies() {
		switch cookie.Name {
		case "accessToken":
			a.AccessToken = cookie.Value
		case "refreshToken":
			a.RefreshToken = cookie.Value
		case "csrfToken":
			a.CsrfToken = cookie.Value
		}
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("login failed: status " + resp.Status)
	}

	return nil
}
