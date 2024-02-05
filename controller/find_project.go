package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fajritsaniy/lpse-screening/usecase"
	"github.com/fajritsaniy/lpse-screening/utils"
)

func FindProjectParticipant(w http.ResponseWriter, sessionID string, searchInput string) {
	auctionList := usecase.AuctionList(sessionID, searchInput)
	// fmt.Println(participantList)

	// Iterate over the slice using range
	for index, value := range auctionList {
		projectWinner := usecase.ProjectWinner(sessionID, value)
		participants := usecase.ParticipantList(sessionID, value)
		auctionList[index].ProjectWinner = projectWinner
		auctionList[index].Participants = participants
	}

	jsonAuctionList, _ := json.Marshal(auctionList)
	// fmt.Println(string(jsonAuctionList))
	fileName := utils.JSONToCSV(w, string(jsonAuctionList), searchInput)
	fmt.Println(fileName, "CSV file has been generated.")
}
