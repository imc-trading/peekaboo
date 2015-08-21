// +build darwin

package mem

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

	i.TotalKB, err = strconv.Atoi(o["hw.memsize"])
	i.TotalKB = i.TotalKB / 1024
	if err != nil {
		return Info{}, err
	}

	return i, nil
}
