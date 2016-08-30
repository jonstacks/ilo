package main

import (
	"fmt"

	"github.com/jonstacks/ilo"
)

func main() {
	fmt.Println("Running the ilo server...")
	ilo.StartServer()
}
