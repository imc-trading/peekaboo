// +build linux

package filesystems

import (
        "os/exec"
	"strings"
	"errors"
	"strconv"
)

type Filesystems []Filesystem

type Filesystem struct {
        Name    string `json:"filesystem"`
        UsedKB   int `json:"usedKB"`
        UsedMB   float32 `json:"usedMB"`
        UsedGB   float32 `json:"usedGB"`
        AvailableKB int      `json:"availableKB"`
        AvailableMB float32      `json:"availableMB"`
        AvailableGB float32      `json:"availableGB"`
	TotalKB	int `json:"totalKB"`
	TotalMB	float32 `json:"totalMB"`
	TotalGB	float32 `json:"totalGB"`
        UsedPct     float32  `json:"usedPct"`
        AvailablePct     float32  `json:"availablePct"`
        MountedOn  string     `json:"mountedOn"`
}

func Get() (Filesystems, error) {
	fs := Filesystems{}

        _, err := exec.LookPath("df")
        if err != nil {
                return Filesystems{}, errors.New("command doesn't exist: df")
        }

        var o []byte

        o, err = exec.Command("df", "-k").Output()
        if err != nil {
                return Filesystems{}, err
        }

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		f := Filesystem{}

		f.Name = v[0]

                f.TotalKB, err = strconv.Atoi(v[1])
                if err != nil {
                        return Filesystems{}, err
                }
                f.TotalMB = float32(f.TotalKB) / 1024
                f.TotalGB = float32(f.TotalMB) / 1024

                f.UsedKB, err = strconv.Atoi(v[2])
                if err != nil {
                        return Filesystems{}, err
                }
                f.UsedMB = float32(f.UsedKB) / 1024
                f.UsedGB = float32(f.UsedMB) / 1024

                f.AvailableKB, err = strconv.Atoi(v[3])
                if err != nil {
                        return Filesystems{}, err
                }
                f.AvailableMB = float32(f.AvailableKB) / 1024
                f.AvailableGB = float32(f.AvailableMB) / 1024

		f.UsedPct = float32(f.UsedKB) / float32(f.TotalKB) * 100
		f.AvailablePct = float32(f.AvailableKB) / float32(f.TotalKB) * 100

		f.MountedOn = v[5]

		fs = append(fs, f)
	}

	return fs, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
