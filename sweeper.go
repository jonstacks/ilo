package ilo

import "net"

type Sweeper struct {
  Subnet string  // Subnet CIDR
}

func NewSweeper(subnet string) (*Sweeper, error) {
  _, _, err := net.ParseCIDR(subnet)
	if err != nil {
    return nil, err
	}
  return &Sweeper{subnet}, nil
}
