package ilo

import (
  "net"
  "sort"
  "time"
)

type Sweeper struct {
  Subnet   string      // Subnet CIDR
  Network  *net.IPNet  // IP mask
}

func NewSweeper(subnet string) (*Sweeper, error) {
  _, ip_net, err := net.ParseCIDR(subnet)
	if err != nil {
    return nil, err
	}
  return &Sweeper{subnet, ip_net}, nil
}

func (s Sweeper) Sweep(timeout time.Duration) []Info {
  var infos []Info
  reply := make(chan *Info)
  quit := time.After(timeout)

  // Spawn all the clients
  for ip := range IPNetwork(s.Network) {
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
    case <- quit:
      break Gather
    case i := <- reply:
      // Got a new ilo result
      infos = append(infos, *i)
    }
  }
  sort.Sort(ByHost(infos))
  return infos
}
