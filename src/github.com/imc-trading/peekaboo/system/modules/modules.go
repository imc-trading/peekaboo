// +build linux

package modules

import (
	"strconv"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type Modules []Module

type Module struct {
	Module string   `json:"module"`
	Size   int      `json:"size"`
	UsedBy []string `json:"usedBy"`
}

func Get() (Modules, error) {
	o, err := parse.Exec("lsmod", []string{})
	if err != nil {
		return Modules{}, err
	}

	m := Modules{}
	first := true
	for _, line := range strings.Split(string(o), "\n") {
		a := strings.Fields(line)
		if len(a) < 3 || first == true {
			first = false
			continue
		}

		size, err := strconv.Atoi(strings.TrimSpace(a[1]))
		if err != nil {
			return Modules{}, err
		}

		usedBy := []string{}
		if len(a) > 3 {
			usedBy = strings.Split(strings.TrimSpace(a[3]), ",")
		}

		m = append(m, Module{
			Module: strings.TrimSpace(a[0]),
			Size:   size,
			UsedBy: usedBy,
		})
	}

	return m, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
