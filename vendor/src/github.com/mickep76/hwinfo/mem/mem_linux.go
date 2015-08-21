// +build linux

package mem

import (
	"github.com/mickep76/hwinfo/common"
	"strconv"
)

// GetInfo return information about a systems memory.
func GetInfo() (Info, error) {
	fields := []string{
		"MemTotal",
	}

	i := Info{}

	o, err := common.LoadFileFields("/proc/meminfo", ":", fields)
	if err != nil {
		return Info{}, err
	}

	i.TotalKB, err = strconv.Atoi(o["MemTotal"])
	if err != nil {
		return Info{}, err
	}

	return i, nil
}
