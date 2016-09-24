package ilo

import (
	"fmt"
	"os"

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

// ByHost implements sort.Interface for []Info
type ByHost []Info

func (h ByHost) Len() int           { return len(h) }
func (h ByHost) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ByHost) Less(i, j int) bool { return h[i].Host < h[j].Host }

// PrintILOTable takes in a list of ILO infos and prints a table for them.
func PrintILOTable(infos []Info) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(50)
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
