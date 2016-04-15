package pcicards

import (
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type PCICards []PCICard

type PCICard struct {
	Slot      string `json:"slot"`
	Class     string `json:"class"`
	Vendor    string `json:"vendor"`
	Device    string `json:"device"`
	SubVendor string `json:"subVendor"`
	SubDevice string `json:"subDevice"`
	Revision  string `json:"revision"`
	ProgIntf  string `json:"progIntf"`
}

func Get() (PCICards, error) {
	o, err := parse.Exec("lspci", []string{"-vmm"})
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
		val := strings.TrimSpace(arr[1])

		switch key {
		case "Slot":
			p.Slot = val
			first = false
		case "Class":
			p.Class = val
		case "Vendor":
			p.Vendor = val
		case "Device":
			p.Device = val
		case "SVendor":
			p.SubVendor = val
		case "SDevice":
			p.SubDevice = val
		case "Rev":
			p.Revision = val
		case "ProgIf":
			p.ProgIntf = val
		}
	}

	return list, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
