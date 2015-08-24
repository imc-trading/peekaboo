// +build darwin

package netinfo

import (
	//	"github.com/mickep76/hwinfo/common"
	//	"fmt"
	"net"
)

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

	/*
		interfaces, _ := net.Interfaces()
		for _, inter := range interfaces {
			fmt.Println(inter.Name, inter.HardwareAddr)
			if addrs, err := inter.Addrs(); err == nil {
				for _, addr := range addrs {
					fmt.Println(inter.Name, "->", addr)
				}
			}
		}
	*/
	return i, nil
}
