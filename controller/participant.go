package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	model "github.com/fajritsaniy/lpse-screening/model"
	utils "github.com/fajritsaniy/lpse-screening/utils"
)

func ParticipantList(sessionID string, project model.Project) {

	fmt.Println("==========================================")
	fmt.Println("Project ID : ", project.ProjectID)
	fmt.Println("Project Name : ", project.ProjectName)
	fmt.Println("Project Source : ", project.ProjectSource)
	fmt.Println("Project Amount : ", project.ProjectAmount)
	fmt.Println("==========================================")
	fmt.Println()
	fmt.Println("Participant: ")
	fmt.Println()

	url := fmt.Sprintf("https://lpse.pu.go.id/eproc4/lelang/%s/peserta", project.ProjectID)

	// Create a map for headers
	headers := map[string]string{
		"authority":  "lpse.pu.go.id",
		"accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"cookie":     sessionID,
		"referer":    "https://lpse.pu.go.id/eproc4/lelang/86330064/pengumumanlelang",
		"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}

	// Create an HTTP request with the payload
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Create an HTTP client
	client := &http.Client{}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the HTML content from the response body
	htmlContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	rows := utils.ExtractTableData(string(htmlContent))
	for n := 0; n < len(rows); n++ {
		if n == 0 {
			continue
		}
		fmt.Println("Nama Peserta : ", utils.RemoveExtraSpaces(rows[n][1]))
		fmt.Println("NPWP : ", rows[n][2])
		fmt.Println("Harga Penawaran : ", rows[n][3])
		fmt.Println("Harga Terkoreksi : ", rows[n][4])
		fmt.Println()
	}
}
