// +build darwin

package os

import (
	"github.com/mickep76/hwinfo/common"
	"runtime"
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

	i.Kernel = runtime.GOOS
	i.Product = o["ProductName"]
	i.ProductVersion = o["ProductVersion"]

	i.KernelVersion, err = common.ExecCmd("uname", []string{"-r"})
	if err != nil {
		return Info{}, err
	}

	return i, nil
}
