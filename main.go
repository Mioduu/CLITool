package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	fmt.Println("Program do trackowania pracy i dobrych nawyków")
	startTimer := time.Now()

	positionTimer := time.NewTimer(time.Minute * 5)
	waterTimer := time.NewTimer(time.Minute * 10)
	walkTimer := time.NewTimer(time.Minute * 30)

	<-positionTimer.C
	beeep.Alert("Baczność!", "Wyprostuj się ;)", "")
	<-waterTimer.C
	beeep.Alert("Woda zdrowia ci doda", "Nawadniaj się ;D", "")
	<-walkTimer.C
	beeep.Alert("Spacer nie zaszkodzi", "Odejdź od biurka i pójdź na spacer", "")

	endTimer := time.Since(startTimer)
	fmt.Printf("Pracujesz od: %v\n sekund", endTimer.Round(time.Second))
	os.Exit(1)
}
