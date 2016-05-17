// +build linux

package sysctls

import (
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type Sysctls []Sysctl

type Sysctl struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Get() (Sysctls, error) {
	sysctls := Sysctls{}

	o, err := parse.Exec("sysctl", []string{"-a"})
	if err != nil {
		return Sysctls{}, err
	}

	for _, line := range strings.Split(o, "\n") {
		vals := strings.Fields(line)
		if len(vals) < 3 {
			continue
		}

		s := Sysctl{}

		s.Key = vals[0]
		s.Value = vals[2]

		sysctls = append(sysctls, s)
	}

	return sysctls, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
