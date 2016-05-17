// +build linux

package routes

import (
	"os/exec"
	"strconv"
	"strings"
)

type Routes []Route

type Route struct {
	Destination string `json:"destination"`
	Gateway     string `json:"gateway"`
	Genmask     string `json:"genmask"`
	Flags       string `json:"flags"`
	MSS         int    `json:"mss"` // Maximum segment size
	Window      int    `json:"window"`
	IRTT        int    `json:"irtt"` // Initial round trip time
	Interface   string `json:"interface"`
}

func Get() (Routes, error) {
	routes := Routes{}

	o, err := exec.Command("netstat", "-rn").Output()
	if err != nil {
		return Routes{}, err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 2 || len(v) < 8 {
			continue
		}

		r := Route{}

		r.Destination = v[0]
		r.Gateway = v[1]
		r.Genmask = v[2]
		r.Flags = v[3]

		r.MSS, err = strconv.Atoi(v[4])
		if err != nil {
			return Routes{}, err
		}

		r.Window, err = strconv.Atoi(v[5])
		if err != nil {
			return Routes{}, err
		}

		r.IRTT, err = strconv.Atoi(v[6])
		if err != nil {
			return Routes{}, err
		}

		r.Interface = v[7]

		routes = append(routes, r)
	}

	return routes, nil
}
