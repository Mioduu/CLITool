package reminders

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

type Reminder struct {
	Name     string
	Interval time.Duration
	Message  string
}

func Start(r Reminder) {
	go func() {
		ticker := time.NewTicker(r.Interval)
		defer ticker.Stop()

		for {
			<-ticker.C
			fmt.Println("Reminder:", r.Name)
			beeep.Alert(r.Name, r.Message, "")
		}

	}()

}
