// +build linux

package mounts

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Mounts []Mount

type Mount struct {
	Source  string `json:"source"`
	Target  string `json:"target"`
	FSType  string `json:"fsType"`
	Options string `json:"options"`
}

func Get() (Mounts, error) {
	mounts := Mounts{}

	fn := "/proc/mounts"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return Mounts{}, fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return Mounts{}, err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		m := Mount{}

		m.Source = v[0]
		m.Target = v[1]
		m.FSType = v[2]
		m.Options = v[3]

		mounts = append(mounts, m)
	}

	return mounts, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
