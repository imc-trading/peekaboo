// +build linux

package memory

import (
	"github.com/mickep76/hwinfo/common"
	"strconv"
	"strings"
)

// Get information about system memory.
func Get() (Memory, error) {
	m := Memory{}

	o, err := common.LoadFileFields("/proc/meminfo", ":", []string{
		"MemTotal",
		"MemFree",
		"MemAvailable",
		"Cached",
		"Committed_AS",
		"HugePages_Total",
		"HugePages_Free",
		"HugePages_Rsvd",
		"HugePages_Surp",
		"Hugepagesize",
	})
	if err != nil {
		return Memory{}, err
	}

	m.TotalGB, err = strconv.Atoi(strings.TrimRight(o["MemTotal"], " kB"))
	if err != nil {
		return Memory{}, err
	}
	m.TotalGB = m.TotalGB / 1024 / 1024

	m.FreeGB, err = strconv.Atoi(strings.TrimRight(o["MemFree"], " kB"))
	if err != nil {
		return Memory{}, err
	}
	m.FreeGB = m.FreeGB / 1024 / 1024

	m.AvailableGB, err = strconv.Atoi(strings.TrimRight(o["MemAvailable"], " kB"))
	if err != nil {
		return Memory{}, err
	}
	m.AvailableGB = m.AvailableGB / 1024 / 1024

	m.CachedGB, err = strconv.Atoi(strings.TrimRight(o["Cached"], " kB"))
	if err != nil {
		return Memory{}, err
	}
	m.CachedGB = m.CachedGB / 1024 / 1024

	m.CommittedActSizeGB, err = strconv.Atoi(strings.TrimRight(o["Committed_AS"], " kB"))
	if err != nil {
		return Memory{}, err
	}
	m.CommittedActSizeGB = m.CommittedActSizeGB / 1024 / 1024

	m.HugePagesTot, err = strconv.Atoi(o["HugePages_Total"])
	if err != nil {
		return Memory{}, err
	}

	m.HugePagesFree, err = strconv.Atoi(o["HugePages_Free"])
	if err != nil {
		return Memory{}, err
	}

	m.HugePagesRsvd, err = strconv.Atoi(o["HugePages_Rsvd"])
	if err != nil {
		return Memory{}, err
	}

	m.HugePageSizeKB, err = strconv.Atoi(strings.TrimRight(o["Hugepagesize"], " kB"))
	if err != nil {
		return Memory{}, err
	}
	m.HugePageSizeKB = m.HugePageSizeKB

	/* Invalidate cache call / force */

	return m, nil
}
