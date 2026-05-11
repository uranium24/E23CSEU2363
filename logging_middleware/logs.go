package loggingmiddleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	
	"github.com/uranium23/E23CSEU2363/pkg/utils"
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
	token := utils.GetAuthToken()
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
