// +build linux

package rpms

import (
	"strconv"
	"strings"
	"time"

	"github.com/imc-trading/peekaboo/parse"
)

type RPMs []RPM

type RPM struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Release   string `json:"release"`
	Arch      string `json:"arch"`
	Installed string `json:"installed"`
}

func Get() (RPMs, error) {
	o, err := parse.Exec("rpm", []string{"-qa", "--queryformat", "%{NAME}|%{VERSION}|%{RELEASE}|%{ARCH}|%{INSTALLTIME}\n"})
	if err != nil {
		return RPMs{}, err
	}

	r := RPMs{}
	for _, line := range strings.Split(string(o), "\n") {
		v := strings.Split(line, "|")
		if len(v) < 5 {
			continue
		}

		sec, err := strconv.ParseInt(v[4], 10, 64)
		if err != nil {
			return RPMs{}, err
		}

		tm := time.Unix(int64(sec), 0)

		r = append(r, RPM{
			Name:      v[0],
			Version:   v[1],
			Release:   v[2],
			Arch:      v[3],
			Installed: tm.Format("2006-01-02T15:04:05Z"),
		})
	}

	return r, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
