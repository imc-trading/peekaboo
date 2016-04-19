package ipmi

import (
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type IPMI struct {
	IpmitoolInstalled bool   `json:"ipmitoolInstalled"`
	IpmitoolVersion   string `json:"ipmitoolVersion"`
}

func Get() (IPMI, error) {
	i := IPMI{}

	// ipmitool
	if err := parse.Exists("ipmitool"); err == nil {
		i.IpmitoolInstalled = true

		o, err := parse.Exec("ipmitool", []string{"-V"})
		if err != nil {
			return IPMI{}, err
		}
		arr := strings.Split(o, " ")
		i.IpmitoolVersion = arr[2]
	} else {
		i.IpmitoolInstalled = false
		return i, nil
	}

	return i, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
