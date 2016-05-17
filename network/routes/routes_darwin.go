// +build darwin

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
	Flags       string `json:"flags"`
	Refs        int    `json:"refs"`
	Use         int    `json:"use"`
	Netif       string `json:"netif"`
	Expire      *int   `json:"expire,omitempty"`
}

func Get() (Routes, error) {
	routes := Routes{}

	o, err := exec.Command("netstat", "-rn").Output()
	if err != nil {
		return nil, err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 4 || len(v) < 6 {
			continue
		}

		r := Route{}

		r.Destination = v[0]
		r.Gateway = v[1]
		r.Flags = v[2]

		r.Refs, err = strconv.Atoi(v[3])
		if err != nil {
			return Routes{}, err
		}

		r.Use, err = strconv.Atoi(v[4])
		if err != nil {
			return Routes{}, err
		}

		r.Netif = v[5]

		if len(v) == 7 {
			v, err := strconv.Atoi(v[6])
			if err != nil {
				return Routes{}, err
			}
			r.Expire = &v
		}

		routes = append(routes, r)
	}

	return routes, nil
}
