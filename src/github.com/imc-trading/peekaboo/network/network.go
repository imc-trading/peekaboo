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
	OnloadInstalled  bool   `json:"onloadInstalled"`
	OnloadVersion    string `json:"onloadVersion"`
	SfctoolInstalled bool   `json:"sfctoolInstalled"`
	SfctoolVersion   string `json:"sfctoolVersion"`
	SfkeyInstalled   bool   `json:"sfkeyInstalled"`
	SfkeyVersion     string `json:"sfkeyVersion"`
}

func Get() (Network, error) {
	n := Network{}

	// ethtool
	n.EthtoolInstalled = false
	if err := parse.Exists("ethtool"); err == nil {
		n.EthtoolInstalled = true

		o, err := parse.Exec("ethtool", []string{"--version"})
		if err != nil {
			return Network{}, err
		}
		arr := strings.Split(o, " ")
		n.EthtoolVersion = arr[2]
	}

	// lldpctl
	n.LldpctlInstalled = false
	if err := parse.Exists("lldpctl"); err == nil {
		n.LldpctlInstalled = true

		o, err := parse.Exec("lldpctl", []string{"-v"})
		if err != nil {
			return Network{}, err
		}
		n.LldpctlVersion = o
	}

	// onload
	n.OnloadInstalled = false
	if err := parse.Exists("onload"); err == nil {
		n.OnloadInstalled = true

		o, err := parse.Exec("onload", []string{"--version"})
		if err != nil {
			return Network{}, err
		}

		for _, line := range strings.Split(o, "\n") {
			arr := strings.SplitN(line, " ", 2)
			if len(arr) < 2 {
				continue
			}

			key := strings.TrimSpace(arr[0])
			val := strings.TrimSpace(arr[1])

			switch key {
			case "OpenOnload":
				n.OnloadVersion = val
			}
		}
	}

	// sfctool
	n.SfctoolInstalled = false
	if err := parse.Exists("sfctool"); err == nil {
		n.SfctoolInstalled = true

		o, err := parse.Exec("sfctool", []string{"--version"})
		if err != nil {
			return Network{}, err
		}
		arr := strings.Split(o, " ")
		n.SfctoolVersion = arr[2]
	}

	// sfckey
	n.SfkeyInstalled = false
	if err := parse.Exists("sfkey"); err == nil {
		n.SfkeyInstalled = true

		o, err := parse.Exec("sfkey", []string{"--version"})
		if err != nil {
			return Network{}, err
		}

		for _, line := range strings.Split(o, "\n") {
			arr := strings.Split(line, ":")
			if len(arr) < 2 {
				continue
			}

			key := strings.TrimSpace(arr[0])
			val := strings.TrimSpace(arr[1])

			switch key {
			case "sfkey firmware update utility":
				n.SfkeyVersion = strings.TrimLeft(val, "v")
			}
		}
	}

	return n, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
