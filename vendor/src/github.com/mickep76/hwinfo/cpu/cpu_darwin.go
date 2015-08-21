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

	i := Info{}

	o, err := common.ExecCmdFields("/usr/sbin/sysctl", []string{"-a"}, ":", fields)
	if err != nil {
		return Info{}, err
	}

	i.CoresPerSocket, err = strconv.Atoi(o["machdep.cpu.core_count"])
	if err != nil {
		return Info{}, err
	}

	i.Physical, err = strconv.Atoi(o["hw.physicalcpu_max"])
	if err != nil {
		return Info{}, err
	}

	i.Logical, err = strconv.Atoi(o["hw.logicalcpu_max"])
	if err != nil {
		return Info{}, err
	}

	i.Sockets = i.Physical / i.CoresPerSocket
	i.ThreadsPerCore = i.Logical / i.Sockets / i.CoresPerSocket
	i.Model = o["machdep.cpu.brand_string"]
	i.Flags = strings.ToLower(o["machdep.cpu.features"])

	return i, nil
}
