package ilo

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// Info is a struct to store ILO Info
type Info struct {
	Host       string
	Serial     string
	Model      string
	ILOVersion string
	Firmware   string
	Success    bool
}

func (i Info) String() string {
	return fmt.Sprintf("%s | %s | %s | %s | %s", i.Host, i.Serial, i.Model, i.ILOVersion, i.Firmware)
}

func cmpIP(ip1, ip2 string) bool {
	octets1 := strings.Split(ip1, ".")
	octets2 := strings.Split(ip2, ".")

	for i := range octets1 {
		o1, _ := strconv.Atoi(octets1[i])
		o2, _ := strconv.Atoi(octets2[i])

		switch {
		case o1 < o2:
			return true
		case o2 < o1:
			return false
		}
	}
	return false
}

// ByHost implements sort.Interface for []Info
type ByHost []Info

func (h ByHost) Len() int           { return len(h) }
func (h ByHost) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ByHost) Less(i, j int) bool { return cmpIP(h[i].Host, h[j].Host) }

// PrintILOTable takes in a list of ILO infos and prints a table for them.
func PrintILOTable(infos []Info) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader(
		[]string{
			"IP",
			"Serial",
			"Model",
			"ILO Version",
			"Firmware Version",
		},
	)

	for _, i := range infos {
		table.Append([]string{
			i.Host,
			i.Serial,
			i.Model,
			i.ILOVersion,
			i.Firmware,
		})
	}
	table.Render()
}
