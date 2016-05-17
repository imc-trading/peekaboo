package pcicards

import (
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type PCICards []PCICard

type PCICard struct {
	Slot        string  `json:"slot"`
	Class       string  `json:"class"`
	ClassId     string  `json:"classId"`
	Vendor      string  `json:"vendor"`
	VendorId    string  `json:"vendorId"`
	Device      string  `json:"device"`
	DeviceId    string  `json:"deviceId"`
	SubVendor   *string `json:"subVendor,omitempty"`
	SubVendorId *string `json:"subVendorId,omitempty"`
	SubDevice   *string `json:"subDevice,omitempty"`
	SubDeviceId *string `json:"subDeviceId,omitempty"`
	Revision    *string `json:"revision,omitempty"`
	ProgIntf    *string `json:"progIntf,omitempty"`
}

func Get() (PCICards, error) {
	o, err := parse.Exec("lspci", []string{"-vmm", "-nn"})
	if err != nil {
		return PCICards{}, err
	}

	p := PCICard{}
	list := PCICards{}
	first := true
	for _, line := range strings.Split(o, "\n") {

		if !first && line == "" {
			list = append(list, p)
			p = PCICard{}
		}

		arr := strings.SplitN(line, ":", 2)
		if len(arr) < 2 {
			continue
		}

		key := strings.TrimSpace(arr[0])

		arr2 := strings.SplitN(arr[1], "[", 2)
		val := strings.TrimSpace(arr2[0])

		id := ""
		if len(arr2) > 1 {
			id = strings.TrimRight(arr2[1], "]")
		}

		switch key {
		case "Slot":
			p.Slot = val
			first = false
		case "Class":
			p.Class = val
			p.ClassId = id
		case "Vendor":
			p.Vendor = val
			p.VendorId = id
		case "Device":
			p.Device = val
			p.DeviceId = id
		case "SVendor":
			if val != "" {
				p.SubVendor = &val
				p.SubVendorId = &id
			}
		case "SDevice":
			if val != "" {
				p.SubDevice = &val
				p.SubDeviceId = &id
			}
		case "Rev":
			if val != "" {
				p.Revision = &val
			}
		case "ProgIf":
			if val != "" {
				p.ProgIntf = &val
			}
		}
	}

	return list, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
