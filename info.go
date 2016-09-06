package ilo

import "fmt"

type Info struct {
  Host        string
  serial      string
  model       string
  ilo_version string
  firmware    string
  success     bool
}

func (i Info) String() string {
  return fmt.Sprintf("%s | %s | %s | %s | %s", i.Host, i.serial, i.model, i.ilo_version, i.firmware)
}

// ByHost implements sort.Interface for []Info
type ByHost []Info

func (h ByHost) Len() int           { return len(h) }
func (h ByHost) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ByHost) Less(i, j int) bool { return h[i].Host < h[j].Host }
