// +build linux

package meminfo

import (
	"github.com/mickep76/hwinfo/common"
	"strconv"
	"strings"
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

	i.TotalGB, err = strconv.Atoi(strings.TrimRight(o["MemTotal"], " kB"))
	if err != nil {
		return Info{}, err
	}
	i.TotalGB = i.TotalGB / 1024 / 1024

	return i, nil
}