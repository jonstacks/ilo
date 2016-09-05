package ilo

import "fmt"

type Info struct {
  host        string
  serial      string
  model       string
  ilo_version string
  firmware    string
  success     bool
}

func (i Info) String() string {
  return fmt.Sprintf("%s | %s | %s | %s | %s", i.host, i.serial, i.model, i.ilo_version, i.firmware)
}
