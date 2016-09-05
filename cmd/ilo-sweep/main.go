package main

import (
	"fmt"
	"os"

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

	ips := []string{"172.18.0.2", "172.18.0.3", "172.18.0.4"}
	for _, ip := range ips{
		c := ilo.NewClient(ip)
		info, _ := c.GetInfo()
		fmt.Println(info)
	}
}
