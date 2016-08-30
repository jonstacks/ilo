package main

import (
	"fmt"
	"github.com/jonstacks/ilo"
	"net/http"
)

const data = `<?xml version="1.0"?>
  <RIMP>
    <HSI>
      <SBSN>CZCxxxx </SBSN>
      <SPN>ProLiant DL380 G5</SPN>
      <UUID>xxxxxxxx</UUID>
      <SP>1</SP>
      <cUUID>0000-0000-0000-0000</cUUID>
      <VIRTUAL>
        <STATE>Inactive</STATE>
        <VID><BSN></BSN>
        <cUUID></cUUID>
        </VID></VIRTUAL>
    </HSI>
    <MP>
      <ST>1</ST>
      <PN>Integrated Lights-Out 2 (iLO 2)</PN>
      <FWRI>2.02</FWRI>
      <BBLK>3; Jul 11 2004</BBLK>
      <HWRI>ASIC:  7</HWRI>
      <SN>00xx00xx00xx      </SN>
      <UUID>ILO000xxx000</UUID>
      <IPM>1</IPM>
      <SSO>0</SSO>
      <PWRM>3.4</PWRM>
    </MP>
  </RIMP>`

func iloData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, data)
}

func main() {
	fmt.Println("Running the ilo server.")
	port := fmt.Sprintf(":%d", ilo.VIRTUAL_MEDIA_PORT)

	finish := make(chan bool)

	mainServer := http.NewServeMux()
	mainServer.HandleFunc("/xmldata", iloData)

	portServer := http.NewServeMux()

	go func() {
		http.ListenAndServe(":80", mainServer)
	}()

	go func() {
		http.ListenAndServe(port, portServer)
	}()

	<-finish
}
