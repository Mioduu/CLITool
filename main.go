package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	fmt.Println("Program do trackowania pracy i dobrych nawyków")
	startTimer := time.Now()

	var wg sync.WaitGroup
	wg.Add(3)

	positionInterval := time.Minute * 1
	waterInterval := time.Minute * 1
	walkInterval := time.Minute * 1

	go func() {
		defer wg.Done()
		start := time.Now()
		for {
			elapsed := time.Since(start)
			remaining := waterInterval - elapsed
			if remaining <= 0 {
				beeep.Alert("Nawadniaj się", "Napij się wody ;D", "")
				break
			}
			fmt.Printf("[Woda] Pozostało: %v\n", remaining.Round(time.Second))
			time.Sleep(time.Second * 5)
		}
	}()

	go func() {
		defer wg.Done()
		start := time.Now()
		for {
			elapsed := time.Since(start)
			remaining := positionInterval - elapsed
			if remaining <= 0 {
				beeep.Alert("Baczność!", "Wyprostuj się ;)", "")
				break
			}
			fmt.Printf("[Pozycja] Pozostało: %v\n", remaining.Round(time.Second))
			time.Sleep(time.Second * 5)
		}
	}()

	go func() {
		defer wg.Done()
		start := time.Now()
		for {
			elapsed := time.Since(start)
			remaining := walkInterval - elapsed
			if remaining <= 0 {
				beeep.Alert("Czas na spacer", "Wyjdź dotknij trawe ;p", "")
				break
			}
			fmt.Printf("[Spacer] Pozostało: %v\n", remaining.Round(time.Second))
			time.Sleep(time.Second * 5)
		}
	}()

	wg.Wait()

	endTimer := time.Since(startTimer)
	fmt.Printf("Pracujesz od: %v minut \n", endTimer.Round(time.Minute))
	os.Exit(1)
}
