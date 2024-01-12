package model

type Participant struct {
	ParticipantName string `json:"participantName"`
	ParticipantNPWP string `json:"participantNPWP"`
	BidPrice        string `json:"bidPrice"`
	CorrectedPrice  string `json:"correctedPrice"`
}
