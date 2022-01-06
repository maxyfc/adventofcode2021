package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Build and runs each day
	for day := 1; day <= 25; day++ {
		fmt.Printf("---- Day %02d ----\n", day)
		main := fmt.Sprintf("day%02d/main.go", day)
		cmd := exec.Command("go", "run", main)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Unable to execute 'go run %s': %s", main, err)
		}
		fmt.Println(string(out))
	}
}
