package scheduler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/uranium23/E23CSEU2363/pkg/utils"
)

func GetTotalMechHoursFromDepots() int {
	godotenv.Load()
	baseUrl := os.Getenv("baseUrl")
	token := utils.GetAuthToken()
	type DepotMD struct {
		ID            int `json:"ID"`
		MechanicHours int `json:"MechanicHours"`
	}
	type DepotInfo struct {
		Depots []DepotMD `json:"depots"`
	}
	url := baseUrl + "/depots"
	req, _ := http.NewRequest("GET", url, http.NoBody)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	dat, _ := io.ReadAll(resp.Body)
	dinfo := DepotInfo{}
	if err := json.Unmarshal(dat, &dinfo); err != nil {
		panic(err)
	}
	totalHours := 0
	for _, depot := range dinfo.Depots {
		totalHours += depot.MechanicHours
	}
	return totalHours
}
