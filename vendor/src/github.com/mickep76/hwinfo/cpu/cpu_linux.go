// +build linux

package cpu

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func (c *cpu) ForceUpdate() error {
	c.cache.LastUpdated = time.Now()
	c.cache.FromCache = false

	if _, err := os.Stat("/proc/cpuinfo"); os.IsNotExist(err) {
		return errors.New("file doesn't exist: /proc/cpuinfo")
	}

	o, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return err
	}

	cpuID := -1
	cpuIDs := make(map[int]bool)
	c.data.CoresPerSocket = 0
	c.data.Logical = 0
	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.Split(line, ":")
		if len(vals) < 1 {
			continue
		}

		v := strings.Trim(strings.Join(vals[1:], " "), " ")
		if c.data.Model == "" && strings.HasPrefix(line, "model name") {
			c.data.Model = v
		} else if c.data.Flags == "" && strings.HasPrefix(line, "flags") {
			c.data.Flags = v
		} else if c.data.CoresPerSocket == 0 && strings.HasPrefix(line, "cpu cores") {
			c.data.CoresPerSocket, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		} else if strings.HasPrefix(line, "processor") {
			c.data.Logical++
		} else if strings.HasPrefix(line, "physical id") {
			cpuID, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
			cpuIDs[cpuID] = true
		}
	}
	c.data.Sockets = int(len(cpuIDs))
	c.data.Physical = c.data.Sockets * c.data.CoresPerSocket
	c.data.ThreadsPerCore = c.data.Logical / c.data.Sockets / c.data.CoresPerSocket

	return nil
}
