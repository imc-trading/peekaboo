// +build linux

package cpu

import (
	"fmt"
	"github.com/mickep76/hwinfo/common"
	"io/ioutil"
	"strconv"
	"strings"
)

func CPUInfo() (CPUInfo, error) {
	c := CPUInfo

	if _, err := os.Stat("/proc/cpuinfo"); os.IsNotExist(err) {
		return []string{}, fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return CPUInfo{}, fmt.Error("can't read file: /proc/cpuinfo")
	}

	cpuID := -1
	cpuIDs := make(map[int]bool)
	c.CoresPerSocket = 0
	c.Logical = 0
	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.Split(line, ":")
		if len(vals) < 1 {
			continue
		}

		v := strings.Trim(strings.Join(vals[1:], " "), " ")
		if c.Model == "" && strings.HasPrefix(line, "model name") {
			c.Model = v
		} else if c.Flags == "" && strings.HasPrefix(line, "flags") {
			c.Flags = v
		} else if c.CoresPerSocket == 0 && strings.HasPrefix(line, "cpu cores") {
			c.CoresPerSocket, err = strconv.ParseInt(v, 10, 0)
			if err != nil {
				return CpuInfo{}, err
			}
		} else if strings.HasPrefix(line, "processor") {
			c.Logical++
		} else if strings.HasPrefix(line, "physical id") {
			cpuID, err = strconv.ParseInt(v, 10, 0)
			if err != nil {
				return CpuInfo{}, err
			}
			cpuIDs[cpuID] = true
		}
	}
	c.Sockets = int(len(cpuIDs))
	c.Physical = c.Sockets * c.CoresPerSocket
	c.ThreadsPerCore = c.Logical / c.Sockets / c.CoresPerSocket

	return c, nil
}
