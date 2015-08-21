// +build linux

package os

import (
	"github.com/mickep76/hwinfo/common"
)

// GetInfo return information about the operating system.
func GetInfo() (Info, error) {
	fields := []string{
		"Distributor ID",
		"Release",
	}

	i := Info{}

	o, err := common.ExecCmdFields("/usr/bin/lsb_release", []string{"-a"}, ":", fields)
	if err != nil {
		return Info{}, err
	}

	i.Product = o["Distributor ID"]
	i.Version = o["Release"]

	return i, nil
}
