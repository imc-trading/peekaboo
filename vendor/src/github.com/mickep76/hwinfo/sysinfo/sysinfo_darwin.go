// +build darwin

package sysinfo

import (
	"github.com/mickep76/hwinfo/common"
)

// GetInfo return information about a systems memory.
func GetInfo() (Info, error) {
	fields := []string{
		"Model Name",
		"Model Identifier",
		"Serial Number",
		"Boot ROM Version",
		"SMC Version",
	}

	i := Info{}
	i.Manufacturer = "Apple Inc."

	o, err := common.ExecCmdFields("/usr/sbin/system_profiler", []string{"SPHardwareDataType"}, ":", fields)
	if err != nil {
		return Info{}, err
	}

	i.Product = o["Model Name"]
	i.ProductVersion = o["Model Identifier"]
	i.SerialNumber = o["Serial Number"]
	i.BootROMVersion = o["Boot ROM Version"]
	i.SMCVersion = o["SMC Version"]

	return i, nil
}
