package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	fmt.Println("Program do trackowania pracy i dobrych nawyków")
	startTimer := time.Now()

	positionInterval := time.Minute * 1
	waterInterval := time.Minute * 1
	walkInterval := time.Minute * 1

	go func() {
		start := time.Now()
		for {
			elapsed := time.Since(start)
			remaining := waterInterval - elapsed
			if remaining <= 0 {
				beeep.Alert("Nawadniaj się", "Napij się wody ;D", "")
				start = time.Now()
			} else {
				fmt.Printf("\n=== [Woda] ===\nPozostało: %v\n", remaining.Round(time.Second))
				time.Sleep(time.Second * 5)
			}
		}
	}()

	go func() {
		start := time.Now()
		for {
			elapsed := time.Since(start)
			remaining := positionInterval - elapsed
			if remaining <= 0 {
				beeep.Alert("Baczność!", "Wyprostuj się ;)", "")
				start = time.Now()
			} else {
				fmt.Printf("\n=== [Pozycja] ===\nPozostało: %v\n", remaining.Round(time.Second))
				time.Sleep(time.Second * 5)
			}
		}
	}()

	go func() {
		start := time.Now()
		for {
			elapsed := time.Since(start)
			remaining := walkInterval - elapsed
			if remaining <= 0 {
				beeep.Alert("Czas na spacer", "Wyjdź dotknij trawe ;p", "")
				start = time.Now()
			} else {
				fmt.Printf("\n=== [Spacer] ===\nPozostało: %v\n", remaining.Round(time.Second))
				time.Sleep(time.Second * 5)
			}
		}
	}()

	for {
		time.Sleep(time.Minute)
		fmt.Printf("Pracujesz od: %v minut \n", time.Since(startTimer).Round(time.Minute))
	}
}
