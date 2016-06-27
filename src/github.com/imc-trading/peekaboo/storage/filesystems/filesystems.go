// +build linux

package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Filesystems []Filesystem

type Filesystem struct {
        Name    `json:"filesystem`
        Used    `json:"used"`
        Available       `json:"available"`
        UsedPerc        `json:"used_perc"`
        MountedOn       `json:"mounted_on"`
}

func Get() (Filesystems, error) {
	fs := Filesystems{}

        _, err := exec.LookPath("df")
        if err != nil {
                return LogVols{}, errors.New("command doesn't exist: df")
        }

        var o []byte

        o, err = exec.Command("df", "-k").Output()
        if err != nil {
                return Filesystems{}, err
        }

/*
Filesystem                          1K-blocks       Used   Available Use% Mounted on
/dev/mapper/vg_root-lv_root          59544948    3541116    53501740   7% /
*/

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		f := Filesystem{}

		m.Name = v[0]
		m.Used = v[1]
		m.Available = v[2]
		m.UsedPerc = v[3]
		m.MountedOn = v[4]

		fs = append(fs, f)
	}

	return fs, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
