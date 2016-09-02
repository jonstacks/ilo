package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	cidr := os.Args[1]

	ip, ip_net, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println(err)
		os.Exit(7)
	}
	fmt.Printf("Sweeping %s for ILO servers\n", ip_net)
	fmt.Printf("IP given: %s\n", ip)
}
