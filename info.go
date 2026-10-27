package ilo

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
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
	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(&iloTableRenderer{renderer.NewBlueprint(tw.Rendition{
			Symbols: tw.NewSymbols(tw.StyleASCII),
		})}),
		tablewriter.WithHeaderAutoWrap(tw.WrapNone),
		tablewriter.WithRowAutoWrap(tw.WrapNone),
		tablewriter.WithTrimSpace(tw.Off),
		tablewriter.WithTrimTab(tw.Off),
		tablewriter.WithTrimLine(tw.Off),
	)
	table.Header(
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

// iloTableRenderer preserves v0.0.5's per-cell numeric alignment.
type iloTableRenderer struct {
	*renderer.Blueprint
}

var iloTableDecimal = regexp.MustCompile(`^-?(?:\d{1,3}(?:,\d{3})*|\d+)(?:\.\d+)?$`)

func (r *iloTableRenderer) Row(row []string, ctx tw.Formatting) {
	for col, cell := range ctx.Row.Current {
		cell.Align = tw.AlignLeft
		if iloTableDecimal.MatchString(strings.TrimSpace(cell.Data)) {
			cell.Align = tw.AlignRight
		}
		ctx.Row.Current[col] = cell
	}
	r.Blueprint.Row(row, ctx)
}

func (r *iloTableRenderer) Line(ctx tw.Formatting) {
	r.Blueprint.Line(ctx)
	// Legacy header-only tables have both a header separator and a bottom border.
	if ctx.Row.Position == tw.Header && ctx.Row.Location == tw.LocationEnd {
		r.Blueprint.Line(ctx)
	}
}
