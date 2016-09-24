package ilo

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const dataFormat = `<?xml version="1.0"?>
  <RIMP>
    <HSI>
      <SBSN>%s</SBSN>
      <SPN>%s</SPN>
      <UUID>xxxxxxxx</UUID>
      <SP>1</SP>
      <cUUID>0000-0000-0000-0000</cUUID>
      <VIRTUAL>
        <STATE>Inactive</STATE>
        <VID>
          <BSN></BSN>
          <cUUID></cUUID>
        </VID>
      </VIRTUAL>
    </HSI>
    <MP>
      <ST>1</ST>
      <PN>%s</PN>
      <FWRI>%s</FWRI>
      <BBLK>3; Jul 11 2004</BBLK>
      <HWRI>ASIC:  7</HWRI>
      <SN>00xx00xx00xx      </SN>
      <UUID>ILO000xxx000</UUID>
      <IPM>1</IPM>
      <SSO>0</SSO>
      <PWRM>3.4</PWRM>
    </MP>
  </RIMP>`

var sProductNames = []string{
	"ProLiant DL380 G5",
	"ProLiant BL480 G6",
	"ProLiant DL360 G9",
	"ProLiant DL580 G9",
}

var productNames = []string{
	"Integrated Lights-Out 2 (iLO 2)",
	"Integrated Lights-Out 3 (iLO 3)",
	"Integrated Lights-Out 4 (iLO 4)",
}
var firmwareVersions = []string{"2.02", "2.04"}

var data string

func iloData(w http.ResponseWriter, r *http.Request) {
	if data == "" {
		rand.Seed(time.Now().UnixNano())
		// Initialize data
		data = fmt.Sprintf(
			dataFormat,
			randomString(10),
			sProductNames[rand.Intn(len(sProductNames))],
			productNames[rand.Intn(len(productNames))],
			firmwareVersions[rand.Intn(len(firmwareVersions))],
		)
	}

	fmt.Fprintf(w, data)
}
