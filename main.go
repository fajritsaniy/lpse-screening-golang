package main

import "github.com/fajritsaniy/lpse-screening/controller"

func main() {
	sessionID := "SPSE_SESSION=971e4b027d0b27d239103a71a3fbff5d5822a74a-___AT=29f5f733dc9a4620c38629d3932003c6966d6302&___TS=1706716799084&___ID=60d3fb1e-8e67-403a-be68-4a6203b380ae"

	controller.FindProjectParticipant(sessionID)
}
