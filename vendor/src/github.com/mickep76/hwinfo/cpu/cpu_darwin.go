// +build darwin

package cpu

import (
	"github.com/mickep76/hwinfo/common"
	"strconv"
	"strings"
)

// GetInfo return information about a systems CPU(s).
func GetInfo() (Info, error) {
	fields := []string{
		"machdep.cpu.core_count",
		"hw.physicalcpu_max",
		"hw.logicalcpu_max",
		"machdep.cpu.brand_string",
		"machdep.cpu.features",
	}

	c := Info{}

	o, err := common.ExecCmdFields("/usr/sbin/sysctl", []string{"-a"}, ":", fields)
	if err != nil {
		return Info{}, err
	}

	c.CoresPerSocket, err = strconv.Atoi(o["machdep.cpu.core_count"])
	if err != nil {
		return Info{}, err
	}

	c.Physical, err = strconv.Atoi(o["hw.physicalcpu_max"])
	if err != nil {
		return Info{}, err
	}

	c.Logical, err = strconv.Atoi(o["hw.logicalcpu_max"])
	if err != nil {
		return Info{}, err
	}

	c.Sockets = c.Physical / c.CoresPerSocket
	c.ThreadsPerCore = c.Logical / c.Sockets / c.CoresPerSocket
	c.Flags = strings.ToLower(o["cpu_flags"])
	c.Model = o["machdep.cpu.brand_string"]

	return c, nil
}
