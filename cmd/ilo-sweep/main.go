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

	sweeper, err := ilo.NewSweeper(cidr)
	handleError(err)

	// Sweep with a timeout of 5 seconds
	infos := sweeper.Sweep(5 * time.Second)
	for _, info := range infos {
		fmt.Println(info)
	}
}
