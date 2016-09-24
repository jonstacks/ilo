package ilo

import (
	"net"
	"sort"
	"time"

	"github.com/jonstacks/goutils/netutils"
)

// Sweeper is a construct for sweeping a subnet for ILO Devices
type Sweeper struct {
	Subnet  string     // Subnet CIDR
	Network *net.IPNet // IP mask
}

// NewSweeper creates & returns a pointer to a new Sweeper
func NewSweeper(subnet string) (*Sweeper, error) {
	_, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}
	return &Sweeper{subnet, ipNet}, nil
}

// Sweep performs the sweep on the Sweeper's subnet
func (s Sweeper) Sweep(timeout time.Duration) []Info {
	var infos []Info
	reply := make(chan *Info)
	quit := time.After(timeout)

	// Spawn all the clients
	for ip := range netutils.IPNetwork(s.Network) {
		go func(x net.IP) {
			c := NewClient(x.String())
			info, err := c.GetInfo()
			if err == nil {
				reply <- info
			}
		}(ip)
	}

Gather:
	for {
		select {
		case <-quit:
			break Gather
		case i := <-reply:
			// Got a new ilo result
			infos = append(infos, *i)
		}
	}
	sort.Sort(ByHost(infos))
	return infos
}
