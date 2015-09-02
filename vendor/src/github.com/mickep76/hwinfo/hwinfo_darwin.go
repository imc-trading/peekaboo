package hwinfo

import (
	"github.com/mickep76/hwinfo/cpu"
	"github.com/mickep76/hwinfo/memory"
	"github.com/mickep76/hwinfo/network"
	"github.com/mickep76/hwinfo/opsys"
	"github.com/mickep76/hwinfo/system"
	"os"
)

// HWInfo information.
type HWInfo struct {
	Hostname string           `json:"hostname"`
	CPU      *cpu.CPU         `json:"cpu"`
	Memory   *memory.Memory   `json:"memory"`
	OpSys    *opsys.OpSys     `json:"opsys"`
	System   *system.System   `json:"system"`
	Network  *network.Network `json:"network"`
}

// Get information about a system.
func Get() (HWInfo, error) {
	i := HWInfo{}

	host, err := os.Hostname()
	if err != nil {
		return HWInfo{}, err
	}
	i.Hostname = host

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

	return i, nil
}
