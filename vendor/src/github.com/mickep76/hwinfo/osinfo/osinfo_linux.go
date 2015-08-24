// +build linux

package osinfo

import (
	"github.com/mickep76/hwinfo/common"
	"runtime"
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

	i.Kernel = runtime.GOOS
	i.Product = o["Distributor ID"]
	i.ProductVersion = o["Release"]

	i.KernelVersion, err = common.ExecCmd("uname", []string{"-r"})
	if err != nil {
		return Info{}, err
	}

	return i, nil
}
