package model

type Project struct {
	ProjectID     string        `json:"projectID"`
	ProjectName   string        `json:"projectName"`
	ProjectSource string        `json:"projectSource"`
	ProjectAmount string        `json:"projectAmount"`
	ProjectWinner string        `json:"projectWinner"`
	Participants  []Participant `json:"participant"`
}
