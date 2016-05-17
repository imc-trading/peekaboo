// +build darwin

package cpu

import (
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

// Get CPU information.
func Get() (CPU, error) {
	c := CPU{}

	m, err := parse.ExecRegexpMap("/usr/sbin/sysctl", []string{"-a"}, ":", "\\S+:\\s\\S+")
	if err != nil {
		return CPU{}, err
	}

	c.CoresPerSocket, err = parse.StrToInt(m, "machdep.cpu.core_count")
	if err != nil {
		return CPU{}, err
	}

	c.Physical, err = parse.StrToInt(m, "hw.physicalcpu_max")
	if err != nil {
		return CPU{}, err
	}

	c.Logical, err = parse.StrToInt(m, "hw.logicalcpu_max")
	if err != nil {
		return CPU{}, err
	}

	c.Sockets = c.Physical / c.CoresPerSocket
	c.ThreadsPerCore = c.Logical / c.Sockets / c.CoresPerSocket
	c.Model = m["machdep.cpu.brand_string"]
	c.Flags = strings.ToLower(m["machdep.cpu.features"])

	return c, nil
}
