package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fajritsaniy/lpse-screening/model"
)

func JSONToCSV(jsonData string) {
	// Unmarshal JSON data into Project struct
	var projects []model.Project
	err := json.Unmarshal([]byte(jsonData), &projects)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a CSV file
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	header := []string{"projectID", "projectName", "projectSource", "projectAmount", "projectWinner"}
	for i := 1; i <= 10; i++ {
		header = append(header, fmt.Sprintf("participant%d", i))
		header = append(header, fmt.Sprintf("participant%d_npwp", i))
		header = append(header, fmt.Sprintf("participant%d_bidPrice", i))
		header = append(header, fmt.Sprintf("participant%d_correctedPrice", i))
	}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing CSV header:", err)
		return
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
			return
		}
	}

	fmt.Println("CSV file created successfully.")
}
