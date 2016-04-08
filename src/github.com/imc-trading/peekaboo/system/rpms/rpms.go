// +build linux

package rpms

import (
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type RPMs []RPM

type RPM struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Release string `json:"release"`
	Arch    string `json:"arch"`
}

func Get() (RPMs, error) {
	o, err := parse.Exec("rpm", []string{"-qa", "--queryformat", "%{NAME}|%{VERSION}|%{RELEASE}|%{ARCH}\n"})
	if err != nil {
		return RPMs{}, err
	}

	r := RPMs{}
	for _, line := range strings.Split(string(o), "\n") {
		v := strings.Split(line, "|")
		if len(v) < 4 {
			continue
		}

		r = append(r, RPM{
			Name:    v[0],
			Version: v[1],
			Release: v[2],
			Arch:    v[3],
		})
	}

	return r, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
