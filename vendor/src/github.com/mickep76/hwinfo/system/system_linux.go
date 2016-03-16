// +build linux

package system

import (
	"time"

	"github.com/mickep76/hwinfo/common"
)

type data struct {
	Manufacturer   string `json:"manufacturer"`
	Product        string `json:"product"`
	ProductVersion string `json:"product_version"`
	SerialNumber   string `json:"serial_number"`
	BIOSVendor     string `json:"bios_vendor"`
	BIOSDate       string `json:"bios_date"`
	BIOSVersion    string `json:"bios_version"`
}

func (s *system) ForceUpdate() error {
	s.cache.LastUpdated = time.Now()
	s.cache.FromCache = false

	o, err := common.LoadFiles([]string{
		"/sys/devices/virtual/dmi/id/chassis_vendor",
		"/sys/devices/virtual/dmi/id/product_name",
		"/sys/devices/virtual/dmi/id/product_version",
		"/sys/devices/virtual/dmi/id/product_serial",
		"/sys/devices/virtual/dmi/id/bios_vendor",
		"/sys/devices/virtual/dmi/id/bios_date",
		"/sys/devices/virtual/dmi/id/bios_version",
	})
	if err != nil {
		return err
	}

	s.data.Manufacturer = o["chassis_vendor"]
	s.data.Product = o["product_name"]
	s.data.ProductVersion = o["product_version"]
	s.data.SerialNumber = o["product_serial"]
	s.data.BIOSVendor = o["bios_vendor"]
	s.data.BIOSDate = o["bios_date"]
	s.data.BIOSVersion = o["bios_version"]

	return nil
}
