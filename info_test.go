package ilo

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmpIP(t *testing.T) {
	assert.True(t, cmpIP("192.168.1.1", "192.168.1.2"))
	assert.True(t, cmpIP("192.168.2.15", "192.168.3.4"))

	assert.False(t, cmpIP("192.1.1.1", "8.8.8.8"))
	assert.False(t, cmpIP("9.9.9.9", "0.255.255.255"))
	assert.False(t, cmpIP("1.1.1.1", "1.1.1.1"), "Same IPs returns false for cmpIP")
}

func TestSortInfos(t *testing.T) {
	infos := []Info{
		Info{"1.2.4.1", "ABC", "Model1", "iLO 4", "2.02", true},
		Info{"1.2.3.10", "DEF", "Model2", "iLO 3", "2.04", true},
		Info{"1.1.1.255", "GHI", "Model3", "iLO 4", "2.02", true},
	}
	sort.Sort(ByHost(infos))

	assert.Equal(t, "1.1.1.255", infos[0].Host, "ByHost correctly sorts Hosts")
	assert.Equal(t, "1.2.3.10", infos[1].Host, "ByHost correctly sorts Hosts")
	assert.Equal(t, "1.2.4.1", infos[2].Host, "ByHost correctly sorts Hosts")
}

func ExamplePrintILOTable() {
	infos := []Info{
		Info{"1.2.3.5", "ABC", "Model1", "iLO 4", "2.02", true},
		Info{"1.2.3.6", "DEF", "Model2", "iLO 3", "2.04", true},
	}
	PrintILOTable(infos)

	// Output:
	// +---------+--------+--------+-------------+------------------+
	// |   IP    | SERIAL | MODEL  | ILO VERSION | FIRMWARE VERSION |
	// +---------+--------+--------+-------------+------------------+
	// | 1.2.3.5 | ABC    | Model1 | iLO 4       |             2.02 |
	// | 1.2.3.6 | DEF    | Model2 | iLO 3       |             2.04 |
	// +---------+--------+--------+-------------+------------------+
}

func ExampleInfo() {
	i := &Info{
		"testhost",
		"ABCDE1234",
		"ProLiant BL480 G6",
		"Integrated Lights-Out 4 (iLO 4)",
		"2.02",
		true,
	}
	fmt.Println(i)
	// Output:
	// testhost | ABCDE1234 | ProLiant BL480 G6 | Integrated Lights-Out 4 (iLO 4) | 2.02
}
