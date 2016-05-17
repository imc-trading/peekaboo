// +build linux

package volgrps

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

type VolGrps []VolGrp

type VolGrp struct {
	Name   string  `json:"name"`
	Attr   string  `json:"attr"`
	SizeB  int     `json:"sizeB"`
	SizeKB float32 `json:"sizeKB"`
	SizeGB float32 `json:"sizeGB"`
	FreeB  int     `json:"freeB"`
	FreeKB float32 `json:"freeKB"`
	FreeGB float32 `json:"freeGB"`
}

func Get() (VolGrps, error) {
	volGrps := VolGrps{}

	_, err := exec.LookPath("vgs")
	if err != nil {
		return VolGrps{}, errors.New("command doesn't exist: vgs")
	}

	var o []byte

	o, err = exec.Command("vgs", "--units", "B", "--readonly").Output()
	if err != nil {
		o, err = exec.Command("vgs", "--units", "B").Output()
		if err != nil {
			return VolGrps{}, err
		}
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		vg := VolGrp{}

		vg.Name = v[0]
		vg.Attr = v[4]

		vg.SizeB, err = strconv.Atoi(strings.TrimRight(v[5], "B"))
		if err != nil {
			return VolGrps{}, err
		}
		vg.SizeKB = float32(vg.SizeB) / 1024
		vg.SizeGB = float32(vg.SizeB) / 1024 / 1024 / 1024

		vg.FreeB, err = strconv.Atoi(strings.TrimRight(v[6], "B"))
		if err != nil {
			return VolGrps{}, err
		}
		vg.FreeKB = float32(vg.FreeB) / 1024
		vg.FreeGB = float32(vg.FreeB) / 1024 / 1024 / 1024

		volGrps = append(volGrps, vg)
	}

	return volGrps, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
