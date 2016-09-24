package ilo

import (
	"fmt"
	"net/http"
)

// StartServer starts a ILO Server that listens for /xmldata
func StartServer() {
	finish := make(chan bool)

	mainServer := http.NewServeMux()
	mainServer.HandleFunc("/xmldata", iloData)

	portServer := http.NewServeMux()

	go func() {
		http.ListenAndServe(":80", mainServer)
	}()

	go func() {
		port := fmt.Sprintf(":%d", VirtualMediaPort)
		http.ListenAndServe(port, portServer)
	}()

	<-finish
}
