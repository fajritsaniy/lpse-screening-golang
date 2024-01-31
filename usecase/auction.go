package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	model "github.com/fajritsaniy/lpse-screening/model"
	utils "github.com/fajritsaniy/lpse-screening/utils"
)

func AuctionList(sessionID string, searchKey string) []model.Project {
	var projects []model.Project
	url := "https://lpse.pu.go.id/eproc4/dt/lelang?rekanan=&kontrak_status=0&instansiId="

	//Authentication Regex Logic
	var authenticationToken string
	re := regexp.MustCompile(`___AT=([^&]+)`)

	// Find the first match
	match := re.FindStringSubmatch(sessionID)

	// Check if a match is found
	if len(match) >= 2 {
		// Extract the captured value
		result := match[1]
		authenticationToken = result
	} else {
		fmt.Println("No match found.")
	}

	offset := "draw=4"
	others := "&columns%5B0%5D%5Bdata%5D=0&columns%5B0%5D%5Bname%5D=&columns%5B0%5D%5Bsearchable%5D=true&columns%5B0%5D%5Borderable%5D=true&columns%5B0%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B0%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B1%5D%5Bdata%5D=1&columns%5B1%5D%5Bname%5D=&columns%5B1%5D%5Bsearchable%5D=true&columns%5B1%5D%5Borderable%5D=true&columns%5B1%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B1%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B2%5D%5Bdata%5D=2&columns%5B2%5D%5Bname%5D=&columns%5B2%5D%5Bsearchable%5D=true&columns%5B2%5D%5Borderable%5D=true&columns%5B2%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B2%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B3%5D%5Bdata%5D=3&columns%5B3%5D%5Bname%5D=&columns%5B3%5D%5Bsearchable%5D=false&columns%5B3%5D%5Borderable%5D=false&columns%5B3%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B3%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B4%5D%5Bdata%5D=4&columns%5B4%5D%5Bname%5D=&columns%5B4%5D%5Bsearchable%5D=true&columns%5B4%5D%5Borderable%5D=true&columns%5B4%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B4%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B5%5D%5Bdata%5D=5&columns%5B5%5D%5Bname%5D=&columns%5B5%5D%5Bsearchable%5D=true&columns%5B5%5D%5Borderable%5D=true&columns%5B5%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B5%5D%5Bsearch%5D%5Bregex%5D=false&order%5B0%5D%5Bcolumn%5D=5&order%5B0%5D%5Bdir%5D=desc&start=0&length=25&search%5Bvalue%5D="
	searchBox := utils.ReplaceSpacesWithPlus(searchKey)
	others1 := "&search%5Bregex%5D=false&authenticityToken="
	payload := offset + others + searchBox + others1 + authenticationToken

	// Create a map for headers
	headers := map[string]string{
		"authority":        "lpse.pu.go.id",
		"accept":           "application/json, text/javascript, */*; q=0.01",
		"content-type":     "application/x-www-form-urlencoded; charset=UTF-8",
		"cookie":           sessionID,
		"x-requested-with": "XMLHttpRequest",
	}

	// Create an HTTP request with the payload
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
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
		return nil
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Please check your Session ID")
		os.Exit(1)
		return nil
	}

	// Access the "data" field in the response
	if data, ok := response["data"].([]interface{}); ok {
		// Iterate through the "data" slice
		for _, entry := range data {
			// Type assertion for each entry in the slice
			if values, ok := entry.([]interface{}); ok {
				// Iterate through the values in the inner slice
				project := model.Project{
					ProjectID:     utils.InterfaceToString(values[0]),
					ProjectName:   utils.InterfaceToString(values[1]),
					ProjectSource: utils.InterfaceToString(values[2]),
					ProjectAmount: utils.InterfaceToString(values[10]),
				}
				projects = append(projects, project)
			}
		}
	} else {
		fmt.Println("Data field not found in the response.")
	}
	return projects
}
