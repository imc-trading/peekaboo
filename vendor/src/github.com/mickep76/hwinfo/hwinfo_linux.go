// +build linux

package hwinfo

import (
	"github.com/mickep76/hwinfo/cpu"
	"github.com/mickep76/hwinfo/disks"
	"github.com/mickep76/hwinfo/lvm"
	"github.com/mickep76/hwinfo/memory"
	"github.com/mickep76/hwinfo/mounts"
	"github.com/mickep76/hwinfo/network"
	"github.com/mickep76/hwinfo/opsys"
	"github.com/mickep76/hwinfo/pci"
	"github.com/mickep76/hwinfo/routes"
	"github.com/mickep76/hwinfo/sysctl"
	"github.com/mickep76/hwinfo/system"
	"os"
	"strings"
)

// HWInfo information.
type HWInfo struct {
	Hostname      string           `json:"hostname"`
	ShortHostname string           `json:"short_hostname"`
	CPU           *cpu.CPU         `json:"cpu"`
	Memory        *memory.Memory   `json:"memory"`
	OpSys         *opsys.OpSys     `json:"opsys"`
	System        *system.System   `json:"system"`
	Network       *network.Network `json:"network"`
	PCI           *[]pci.PCI       `json:"pci"`
	Disks         *[]disks.Disk    `json:"disks"`
	Routes        *[]routes.Route  `json:"routes"`
	Sysctl        *[]sysctl.Sysctl `json:"sysctl"`
	LVM           *lvm.LVM         `json:"lvm"`
	Mounts        *[]mounts.Mount  `json:"mounts"`
}

// Get information about a system.
func Get() (HWInfo, error) {
	i := HWInfo{}

	host, err := os.Hostname()
	if err != nil {
		return HWInfo{}, err
	}
	i.Hostname = host
	i.ShortHostname = strings.Split(host, ".")[0]

	i2, err := cpu.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.CPU = &i2

	i3, err := memory.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.Memory = &i3

	i4, err := opsys.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.OpSys = &i4

	i5, err := system.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.System = &i5

	i6, err := network.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.Network = &i6

	i7, err := pci.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.PCI = &i7

	i8, err := disks.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.Disks = &i8

	i9, err := routes.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.Routes = &i9

	i10, err := sysctl.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.Sysctl = &i10

	i11, err := lvm.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.LVM = &i11

	i12, err := mounts.Get()
	if err != nil {
		return HWInfo{}, err
	}
	i.Mounts = &i12

	return i, nil
}
