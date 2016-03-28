// +build darwin

package memory

import (
	"github.com/imc-trading/peekaboo/parse"
)

type Memory struct {
	TotalKB int `json:"totalKB"`
	TotalGB int `json:"totalGB"`
}

func Get() (Memory, error) {
	mem := Memory{}

	m, err := parse.ExecRegexpMap("/usr/sbin/sysctl", []string{"-a"}, ":", "\\S+:\\s\\S+")
	if err != nil {
		return Memory{}, err
	}

	mem.TotalKB, err = parse.StrToInt(m, "hw.memsize")
	if err != nil {
		return Memory{}, err
	}
	mem.TotalKB = mem.TotalKB / 1024
	mem.TotalGB = mem.TotalKB / 1024 / 1024

	return mem, nil
}
