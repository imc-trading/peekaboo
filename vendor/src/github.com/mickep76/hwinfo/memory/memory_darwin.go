// +build darwin

package memory

import (
	"strconv"
	"time"

	"github.com/mickep76/hwinfo/common"
)

type Data struct {
	TotalKB int `json:"total_kb"`
	TotalGB int `json:"total_gb"`
}

func (m *memory) ForceUpdate() error {
	m.cache.LastUpdated = time.Now()
	m.cache.FromCache = false

	o, err := common.ExecCmdFields("/usr/sbin/sysctl", []string{"-a"}, ":", []string{
		"hw.memsize",
	})
	if err != nil {
		return err
	}

	m.data.TotalKB, err = strconv.Atoi(o["hw.memsize"])
	if err != nil {
		return err
	}
	m.data.TotalKB = m.data.TotalKB / 1024
	m.data.TotalGB = m.data.TotalKB / 1024 / 1024

	return nil
}
