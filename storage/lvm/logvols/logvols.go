// +build linux

package logvols

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

type LogVols []LogVol

type LogVol struct {
	Name   string  `json:"name"`
	VolGrp string  `json:"volGrp"`
	Attr   string  `json:"attr"`
	SizeB  int     `json:"sizeB"`
	SizeKB float32 `json:"sizeKB"`
	SizeGB float32 `json:"sizeGB"`
}

func Get() (LogVols, error) {
	logVols := LogVols{}

	_, err := exec.LookPath("lvs")
	if err != nil {
		return LogVols{}, errors.New("command doesn't exist: lvs")
	}

	var o []byte

	o, err = exec.Command("lvs", "--units", "B", "--readonly").Output()
	if err != nil {
		o, err = exec.Command("lvs", "--units", "B").Output()
		if err != nil {
			return LogVols{}, err
		}
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		l := LogVol{}

		l.Name = v[0]
		l.VolGrp = v[1]
		l.Attr = v[2]

		l.SizeB, err = strconv.Atoi(strings.TrimRight(v[3], "B"))
		if err != nil {
			return LogVols{}, err
		}
		l.SizeKB = float32(l.SizeB) / 1024
		l.SizeGB = float32(l.SizeB) / 1024 / 1024 / 1024

		logVols = append(logVols, l)
	}

	return logVols, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
