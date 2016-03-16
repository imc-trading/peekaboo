// +build darwin

package system

import (
	"time"

	"github.com/mickep76/hwinfo/common"
)

type Data struct {
	Manufacturer   string `json:"manufacturer"`
	Product        string `json:"product"`
	ProductVersion string `json:"product_version"`
	SerialNumber   string `json:"serial_number"`
	BootROMVersion string `json:"boot_rom_version"`
	SMCVersion     string `json:"smc_version"`
}

func (s *system) ForceUpdate() error {
	s.cache.LastUpdated = time.Now()
	s.cache.FromCache = false
	s.data.Manufacturer = "Apple Inc."

	o, err := common.ExecCmdFields("/usr/sbin/system_profiler", []string{"SPHardwareDataType"}, ":", []string{
		"Model Name",
		"Model Identifier",
		"Serial Number",
		"Boot ROM Version",
		"SMC Version",
	})
	if err != nil {
		return err
	}

	s.data.Product = o["Model Name"]
	s.data.ProductVersion = o["Model Identifier"]
	s.data.SerialNumber = o["Serial Number"]
	s.data.BootROMVersion = o["Boot ROM Version"]
	s.data.SMCVersion = o["SMC Version"]

	return nil
}
