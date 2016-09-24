package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jonstacks/ilo"
)

func handleError(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	cidr := os.Args[1]
	timeout := 2 * time.Second

	sweeper, err := ilo.NewSweeper(cidr)
	handleError(err)

	// Sweep with a timeout of 5 seconds
	infos := sweeper.Sweep(timeout)
	ilo.PrintILOTable(infos)
}
