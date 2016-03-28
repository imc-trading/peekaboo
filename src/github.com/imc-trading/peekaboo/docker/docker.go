package docker

import (
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type Docker struct {
	Version string `json:"version"`
	Build   string `json:"build"`
	Running bool   `json:"running"`
}

func Get() (Docker, error) {
	d := Docker{}

	o, err := parse.Exec("docker", []string{"--version"})
	if err != nil {
		return Docker{}, err
	}

	v := strings.Fields(o)

	d.Version = strings.TrimRight(v[2], ",")
	d.Build = v[4]

	_, err2 := parse.Exec("docker", []string{"ps"})
	if err2 != nil {
		d.Running = false
	} else {
		d.Running = true
	}

	return d, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
