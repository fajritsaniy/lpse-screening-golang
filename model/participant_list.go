package model

type Participant struct {
	Number          int    `json:"number"`
	ParticipantName string `json:"participantName"`
	ParticipantNPWP string `json:"participantNPWP"`
	BidPrice        string `json:"bidPrice"`
	CorrectedPrice  string `json:"correctedPrice"`
}

func NewParticipant(number int, name, npwp, bidPrice, correctedPrice string) *Participant {
	return &Participant{
		Number:          number,
		ParticipantName: name,
		ParticipantNPWP: npwp,
		BidPrice:        bidPrice,
		CorrectedPrice:  correctedPrice,
	}
}
