package loggingmiddleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Logger(stack, level, pkg, message string) {
	logBody := struct {
		Stack   string `json:"stack"`
		Level   string `json:"level"`
		Package string `json:"package"`
		Message string `json:"message"`
	}{
		stack, level, pkg, message,
	}
	json, err := json.Marshal(logBody)
	if err != nil {
		panic(err)
	}
	makeReq(json, "http://4.224.186.213/evaluation-service/logs")
}

func makeReq(jsonBody []byte, url string) {
	godotenv.Load()
	token := GetAuthToken()
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	logResp := struct {
		LogID string `json:"logID"`
		Msg   string `json:"message"`
	}{}
	data, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &logResp); err != nil {
		panic(err)
	}
	fmt.Println(logResp)
}

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
