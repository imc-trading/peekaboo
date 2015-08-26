// +build linux

package sysinfo

import (
	"github.com/mickep76/hwinfo/common"
)

// GetInfo return information about a systems memory.
func GetInfo() (Info, error) {
	i := Info{}

	files := []string{
		"/sys/devices/virtual/dmi/id/chassis_vendor",
		"/sys/devices/virtual/dmi/id/product_name",
		"/sys/devices/virtual/dmi/id/product_version",
		"/sys/devices/virtual/dmi/id/product_serial",
		"/sys/devices/virtual/dmi/id/bios_vendor",
		"/sys/devices/virtual/dmi/id/bios_date",
		"/sys/devices/virtual/dmi/id/bios_version",
	}

	o, err := common.LoadFiles(files)
	if err != nil {
		return Info{}, err
	}

	i.Manufacturer = o["chassis_vendor"]
	i.Product = o["product_name"]
	i.ProductVersion = o["product_version"]
	i.SerialNumber = o["product_serial"]
	i.BIOSVendor = o["bios_vendor"]
	i.BIOSDate = o["bios_date"]
	i.BIOSVersion = o["bios_version"]

	return i, nil
}
