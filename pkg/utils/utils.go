package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetAuthToken() (token string) {
	url := "http://4.224.186.213/evaluation-service/auth"
	godotenv.Load()
	body := struct {
		Email        string `json:"email"`
		Name         string `json:"name"`
		RollNo       string `json:"rollNo"`
		AccessCode   string `json:"accessCode"`
		ClientID     string `json:"clientID"`
		ClientSecret string `json:"clientSecret"`
	}{
		Email:        os.Getenv("email"),
		Name:         os.Getenv("name"),
		RollNo:       os.Getenv("rollNo"),
		AccessCode:   os.Getenv("accessCode"),
		ClientID:     os.Getenv("clientID"),
		ClientSecret: os.Getenv("clientSecret"),
	}
	jb, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jb))
	client := &http.Client{}
	resp, _ := client.Do(req)
	rb := struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}{}
	data, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &rb); err != nil {
		panic(err)
	}
	return rb.AccessToken
}
