package hwinfo

import (
	"github.com/mickep76/hwinfo/cpu"
	"github.com/mickep76/hwinfo/mem"
	"github.com/mickep76/hwinfo/netinfo"
	hwos "github.com/mickep76/hwinfo/os"
	"github.com/mickep76/hwinfo/sys"
	"os"
)

// Info structure for information a system.
type Info struct {
	Hostname string          `json:"hostname"`
	CPU      *cpu.Info       `json:"cpu"`
	Mem      *mem.Info       `json:"memory"`
	OS       *hwos.Info      `json:"os"`
	Sys      *sys.Info       `json:"system"`
	Net      *[]netinfo.Info `json:"network"`
}

// GetInfo return information about a system.
func GetInfo() (Info, error) {
	h := Info{}

	host, err := os.Hostname()
	if err != nil {
		return Info{}, err
	}
	h.Hostname = host

	c, err := cpu.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.CPU = &c

	m, err := mem.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.Mem = &m

	o, err := hwos.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.OS = &o

	s, err := sys.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.Sys = &s

	n, err := netinfo.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.Net = &n

	return h, nil
}
