package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fajritsaniy/lpse-screening/controller"
)

func main() {
	sessionID := "SPSE_SESSION=2bad74c68a5316f44809dd322d12f90d047e31c0-___AT=e13a2a572dc4bc093f2cc741493baaa0aa4a0e06&___TS=1705050809477&___ID=10b0ca93-f2ab-46ae-b8a9-194317356778"
	fmt.Print("Masukkan Keyword: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	auctionList := controller.AuctionList(sessionID, scanner.Text())
	// fmt.Println(participantList)

	// Iterate over the slice using range
	for _, value := range auctionList {
		controller.ParticipantList(sessionID, value)
		fmt.Println("=================BREAK=================")
	}
}
