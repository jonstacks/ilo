package ilo

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "net/http"

  "gopkg.in/xmlpath.v2"
)

type Client struct {
  host string
}

var paths = map[string]*xmlpath.Path{
  "SBSN": xmlpath.MustCompile("/RIMP/HSI/SBSN"),
  "SPN":  xmlpath.MustCompile("/RIMP/HSI/SPN"),
  "UUID": xmlpath.MustCompile("/RIMP/HSI/UUID"),
  "PN":   xmlpath.MustCompile("/RIMP/MP/PN"),
  "FWRI": xmlpath.MustCompile("/RIMP/MP/FWRI"),
}

func NewClient(host string) Client {
  return Client{host}
}

func (c Client) GetInfo() (*Info, error) {
  url := fmt.Sprintf("http://%s/xmldata?item=all", c.host)
  resp, err := http.Get(url)
  if err != nil {
    return &Info{c.host, "", "", "", "", false}, err
  }

  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)

  root, err := xmlpath.Parse(bytes.NewReader(body))
  if err != nil {
    return nil, err
  }

  data := make(map[string]string)

  for k := range paths {
    if value, ok := paths[k].String(root); ok {
      data[k] = value
    }
  }

  return &Info{c.host, data["SBSN"], data["SPN"], data["PN"], data["FWRI"], true}, nil
}
