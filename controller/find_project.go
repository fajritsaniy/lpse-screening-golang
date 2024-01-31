package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fajritsaniy/lpse-screening/usecase"
	"github.com/fajritsaniy/lpse-screening/utils"
)

func FindProjectParticipant(sessionID string) {
	fmt.Print("Masukkan Keyword: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	auctionList := usecase.AuctionList(sessionID, scanner.Text())
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
	utils.JSONToCSV(string(jsonAuctionList))
}
