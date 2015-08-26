package hwinfo

import (
	"github.com/mickep76/hwinfo/cpuinfo"
	"github.com/mickep76/hwinfo/meminfo"
	"github.com/mickep76/hwinfo/netinfo"
	"github.com/mickep76/hwinfo/osinfo"
	"github.com/mickep76/hwinfo/pciinfo"
	"github.com/mickep76/hwinfo/sysinfo"
	"os"
)

// Info structure for information a system.
type Info struct {
	Hostname string        `json:"hostname"`
	CPU      *cpuinfo.Info `json:"cpu"`
	Mem      *meminfo.Info `json:"memory"`
	OS       *osinfo.Info  `json:"os"`
	Sys      *sysinfo.Info `json:"system"`
	Net      *netinfo.Info `json:"network"`
	PCI      *pciinfo.Info `json:"pci,omitempty"`
}

// GetInfo return information about a system.
func GetInfo() (Info, error) {
	h := Info{}

	host, err := os.Hostname()
	if err != nil {
		return Info{}, err
	}
	h.Hostname = host

	c, err := cpuinfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.CPU = &c

	m, err := meminfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.Mem = &m

	o, err := osinfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.OS = &o

	s, err := sysinfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.Sys = &s

	n, err := netinfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.Net = &n

	p, err := pciinfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.PCI = &p

	return h, nil
}
