// +build linux

package arp

import (
	"os/exec"
	"strings"
)

type ArpEntries []Arp

type Arp struct {
	Address string `json:"address"`
	HWtype     string `json:"hwtype"`
	HWaddress     string `json:"hwaddress"`
	Flags       string `json:"flags"`
	Interface   string `json:"interface"`
}

func Get() (ArpEntries, error) {
	arpentries := ArpEntries{}

	o, err := exec.Command("arp", "-an").Output()
	if err != nil {
		return ArpEntries{}, err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 2 || len(v) < 8 {
			continue
		}

		a := Arp{}

		a.Address = v[0]
		a.HWtype = v[1]
		a.HWaddress = v[2]
		a.Flags = v[3]
		a.Interface = v[4]

		arpentries = append(arpentries, a)
	}

	return arpentries, nil
}
