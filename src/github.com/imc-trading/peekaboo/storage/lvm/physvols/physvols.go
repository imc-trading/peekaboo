// +build linux

package physvols

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

type PhysVols []PhysVol

type PhysVol struct {
	Name   string  `json:"name"`
	VolGrp string  `json:"volGrp"`
	Format string  `json:"format"`
	Attr   string  `json:"attr"`
	SizeB  int     `json:"sizeB"`
	SizeKB float32 `json:"sizeKB"`
	SizeGB float32 `json:"sizeGB"`
	FreeB  int     `json:"freeB"`
	FreeKB float32 `json:"freeKB"`
	FreeGB float32 `json:"freeGB"`
}

func Get() (PhysVols, error) {
	physVols := PhysVols{}

	_, err := exec.LookPath("pvs")
	if err != nil {
		return PhysVols{}, errors.New("command doesn't exist: pvs")
	}

	var o []byte

	o, err = exec.Command("pvs", "--units", "B", "--readonly").Output()
	if err != nil {
		o, err = exec.Command("pvs", "--units", "B").Output()
		if err != nil {
			return PhysVols{}, err
		}
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		p := PhysVol{}

		p.Name = v[0]
		p.VolGrp = v[1]
		p.Format = v[2]
		p.Attr = v[3]

		p.SizeB, err = strconv.Atoi(strings.TrimRight(v[4], "B"))
		if err != nil {
			return PhysVols{}, err
		}
		p.SizeKB = float32(p.SizeB) / 1024
		p.SizeGB = float32(p.SizeB) / 1024 / 1024 / 1024

		p.FreeB, err = strconv.Atoi(strings.TrimRight(v[5], "B"))
		if err != nil {
			return PhysVols{}, err
		}
		p.FreeKB = float32(p.FreeB) / 1024
		p.FreeGB = float32(p.FreeB) / 1024 / 1024 / 1024

		physVols = append(physVols, p)
	}

	return physVols, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
