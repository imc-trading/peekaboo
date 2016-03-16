// +build darwin

package cpu

import (
	"strconv"
	"strings"
	"time"

	"github.com/mickep76/hwinfo/common"
)

func (c *cpu) ForceUpdate() error {
	c.cache.LastUpdated = time.Now()
	c.cache.FromCache = false

	o, err := common.ExecCmdFields("/usr/sbin/sysctl", []string{"-a"}, ":", []string{
		"machdep.cpu.core_count",
		"hw.physicalcpu_max",
		"hw.logicalcpu_max",
		"machdep.cpu.brand_string",
		"machdep.cpu.features",
	})
	if err != nil {
		return err
	}

	c.data.CoresPerSocket, err = strconv.Atoi(o["machdep.cpu.core_count"])
	if err != nil {
		return err
	}

	c.data.Physical, err = strconv.Atoi(o["hw.physicalcpu_max"])
	if err != nil {
		return err
	}

	c.data.Logical, err = strconv.Atoi(o["hw.logicalcpu_max"])
	if err != nil {
		return err
	}

	c.data.Sockets = c.data.Physical / c.data.CoresPerSocket
	c.data.ThreadsPerCore = c.data.Logical / c.data.Sockets / c.data.CoresPerSocket
	c.data.Model = o["machdep.cpu.brand_string"]
	c.data.Flags = strings.ToLower(o["machdep.cpu.features"])

	return nil
}
