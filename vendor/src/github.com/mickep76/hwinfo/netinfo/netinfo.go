package netinfo

import (
	"net"
)

// Info structure for information about a systems memory.
type Info struct {
	Name   string   `json:"name"`
	MTU    int      `json:"mtu"`
	IPAddr []string `json:"ipaddr"`
	HWAddr string   `json:"hwaddr"`
}

// GetInfo return information about a systems memory.
func GetInfo() ([]Info, error) {
	i := []Info{}

	ifs, err := net.Interfaces()
	if err != nil {
		return []Info{}, err
	}

	for _, v := range ifs {
		if v.Name == "lo" || v.Name == "lo0" {
			continue
		}

		addrs, err := v.Addrs()
		if err != nil {
			return []Info{}, err
		}

		ia := []string{}
		for _, addr := range addrs {
			ia = append(ia, addr.String())
		}

		i = append(i, Info{
			Name:   v.Name,
			HWAddr: v.HardwareAddr.String(),
			MTU:    v.MTU,
			IPAddr: ia,
		})
	}

	return i, nil
}
