package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fajritsaniy/lpse-screening/model"
)

func JSONToCSV(w http.ResponseWriter, jsonData string, searchInput string) string {
	currentTime := time.Now()

	// Format the current date as a string in "2006-01-02" format (Year, Month, Day)
	currentDateString := currentTime.Format("2006-01-02 15:04:05")

	fileName := fmt.Sprintf("%s - Project List - %s", searchInput, currentDateString)
	contentDisposition := "attachment;filename=" + fileName + ".csv"

	// Set Content-Type header to text/csv
	w.Header().Set("Content-Type", "text/csv")
	// Set Content-Disposition header to trigger download with the filename "empty.csv"
	w.Header().Set("Content-Disposition", contentDisposition)

	// Unmarshal JSON data into Project struct
	var projects []model.Project
	err := json.Unmarshal([]byte(jsonData), &projects)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// Create a CSV writer
	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Write header row
	header := []string{"ID Proyek", "Nama", "Sumber Proyek", "Total Proyek", "Pemenang"}
	for i := 1; i <= 10; i++ {
		header = append(header, fmt.Sprintf("Peserta %d", i))
		header = append(header, fmt.Sprintf("NPWP Peserta %d", i))
		header = append(header, fmt.Sprintf("Harga Penawaran Peserta %d", i))
		header = append(header, fmt.Sprintf("Harga Terkoreksi Peserta %d", i))
	}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing CSV header:", err)
		return ""
	}

	// Write project and participant data for each project
	for _, project := range projects {
		projectRow := []string{project.ProjectID, project.ProjectName, project.ProjectSource, project.ProjectAmount, project.ProjectWinner}
		for _, participant := range project.Participants {
			projectRow = append(projectRow, participant.ParticipantName, participant.ParticipantNPWP, participant.BidPrice, participant.CorrectedPrice)
		}
		err := writer.Write(projectRow)
		if err != nil {
			fmt.Println("Error writing project and participant data:", err)
			return ""
		}
	}
	return fileName
}
