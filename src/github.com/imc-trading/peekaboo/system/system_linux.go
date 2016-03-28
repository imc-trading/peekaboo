// +build linux

package system

import (
	"os"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
	"github.com/imc-trading/peekaboo/version"
)

type System struct {
	Hostname        string `json:"hostname"`
	ShortHostname   string `json:"shortHostname"`
	PeekabooVersion string `json:"peekabooVersion"`
	Manufacturer    string `json:"manufacturer"`
	Product         string `json:"product"`
	ProductVersion  string `json:"productVersion"`
	SerialNumber    string `json:"serialNumber"`
	BIOSVendor      string `json:"biosVendor"`
	BIOSDate        string `json:"biosDate"`
	BIOSVersion     string `json:"biosVersion"`
}

func Get() (System, error) {
	s := System{}

	host, err := os.Hostname()
	if err != nil {
		return System{}, err
	}
	s.Hostname = host
	s.ShortHostname = strings.Split(host, ".")[0]

	s.PeekabooVersion = version.Version

	m, err := parse.LoadFiles([]string{
		"/sys/devices/virtual/dmi/id/sys_vendor",
		"/sys/devices/virtual/dmi/id/product_name",
		"/sys/devices/virtual/dmi/id/product_version",
		"/sys/devices/virtual/dmi/id/product_serial",
		"/sys/devices/virtual/dmi/id/bios_vendor",
		"/sys/devices/virtual/dmi/id/bios_date",
		"/sys/devices/virtual/dmi/id/bios_version",
	})
	if err != nil {
		return System{}, err
	}

	s.Manufacturer = m["sys_vendor"]
	s.Product = m["product_name"]
	s.ProductVersion = m["product_version"]
	s.SerialNumber = m["product_serial"]
	s.BIOSVendor = m["bios_vendor"]
	s.BIOSDate = m["bios_date"]
	s.BIOSVersion = m["bios_version"]

	return s, nil
}
