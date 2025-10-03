package main

import (
	"CLITOOL/ui"
	"fmt"
	"os"
)

func main() {
	finalModel, err := ui.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(finalModel.Selected) == 0 {
		fmt.Println("Nothing selected - Nothing to start")
	}
}
