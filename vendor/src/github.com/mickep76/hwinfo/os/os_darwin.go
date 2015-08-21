// +build darwin

package os

import (
	"github.com/mickep76/hwinfo/common"
)

// GetInfo return information about the operating system.
func GetInfo() (Info, error) {
	fields := []string{
		"ProductName",
		"ProductVersion",
	}

	i := Info{}

	o, err := common.ExecCmdFields("/usr/bin/sw_vers", []string{}, ":", fields)
	if err != nil {
		return Info{}, err
	}

	i.Product = o["ProductName"]
	i.Version = o["ProductVersion"]

	return i, nil
}
