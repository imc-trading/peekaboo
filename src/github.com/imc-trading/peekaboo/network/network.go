package network

import (
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type Network struct {
	EthtoolInstalled bool   `json:"ethtoolInstalled"`
	EthtoolVersion   string `json:"ethtoolVersion"`
	LldpctlInstalled bool   `json:"lldpctlInstalled"`
	LldpctlVersion   string `json:"lldpctlVersion"`
}

func Get() (Network, error) {
	n := Network{}

	n.EthtoolInstalled = false
	if err := parse.Exists("ethtool"); err == nil {
		n.EthtoolInstalled = true
	}

	n.LldpctlInstalled = false
	if err := parse.Exists("lldpctl"); err == nil {
		n.LldpctlInstalled = true
	}

	o, err := parse.Exec("ethtool", []string{"--version"})
	if err != nil {
		return Network{}, err
	}
	arr := strings.Split(o, " ")
	n.EthtoolVersion = arr[2]

	o2, err := parse.Exec("lldpctl", []string{"-v"})
	if err != nil {
		return Network{}, err
	}
	n.LldpctlVersion = o2

	return n, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
