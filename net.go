package ilo

import "net"

func increment_ip(ip net.IP) {
  for j := len(ip)-1; j>=0; j-- {
    ip[j]++
    if ip[j] > 0 {
      break
    }
  }
}

func first_address(n *net.IPNet) net.IP {
  return copy_address(n.IP)
}

func copy_address(a net.IP) net.IP {
  x := make(net.IP, len(a))
  copy(x, a)
  return x
}

func IPNetwork(network *net.IPNet) <-chan net.IP {
  out := make(chan net.IP)

  go func() {
    for ip := first_address(network); network.Contains(ip); increment_ip(ip) {
      out <- copy_address(ip)
    }
    close(out)
  }()
  return out
}
