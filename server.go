package ilo

import (
  "fmt"
  "net/http"
)

func StartServer() {
  finish := make(chan bool)

	mainServer := http.NewServeMux()
	mainServer.HandleFunc("/xmldata", iloData)

	portServer := http.NewServeMux()

	go func() {
		http.ListenAndServe(":80", mainServer)
	}()

	go func() {
    port := fmt.Sprintf(":%d", VIRTUAL_MEDIA_PORT)
		http.ListenAndServe(port, portServer)
	}()

	<-finish
}
