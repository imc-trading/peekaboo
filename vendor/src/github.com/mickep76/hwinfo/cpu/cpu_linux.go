// +build linux

package cpu

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// GetInfo return information about a systems CPU(s).
func GetInfo() (Info, error) {
	i := Info{}

	if _, err := os.Stat("/proc/cpuinfo"); os.IsNotExist(err) {
		return Info{}, errors.New("file doesn't exist: /proc/cpuinfo")
	}

	o, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return Info{}, err
	}

	cpuID := -1
	cpuIDs := make(map[int]bool)
	i.CoresPerSocket = 0
	i.Logical = 0
	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.Split(line, ":")
		if len(vals) < 1 {
			continue
		}

		v := strings.Trim(strings.Join(vals[1:], " "), " ")
		if i.Model == "" && strings.HasPrefix(line, "model name") {
			i.Model = v
		} else if i.Flags == "" && strings.HasPrefix(line, "flags") {
			i.Flags = v
		} else if i.CoresPerSocket == 0 && strings.HasPrefix(line, "cpu cores") {
			i.CoresPerSocket, err = strconv.Atoi(v)
			if err != nil {
				return Info{}, err
			}
		} else if strings.HasPrefix(line, "processor") {
			i.Logical++
		} else if strings.HasPrefix(line, "physical id") {
			cpuID, err = strconv.Atoi(v)
			if err != nil {
				return Info{}, err
			}
			cpuIDs[cpuID] = true
		}
	}
	i.Sockets = int(len(cpuIDs))
	i.Physical = i.Sockets * i.CoresPerSocket
	i.ThreadsPerCore = i.Logical / i.Sockets / i.CoresPerSocket

	return i, nil
}
