// +build darwin

package meminfo

import (
	"github.com/mickep76/hwinfo/common"
	"strconv"
)

// GetInfo return information about a systems memory.
func GetInfo() (Info, error) {
	fields := []string{
		"hw.memsize",
	}

	i := Info{}

	o, err := common.ExecCmdFields("/usr/sbin/sysctl", []string{"-a"}, ":", fields)
	if err != nil {
		return Info{}, err
	}

	i.TotalGB, err = strconv.Atoi(o["hw.memsize"])
	if err != nil {
		return Info{}, err
	}
	i.TotalGB = i.TotalGB / 1024 / 1024 / 1024

	return i, nil
}
