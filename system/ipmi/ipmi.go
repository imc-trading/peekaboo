package ipmi

import (
	"net"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type IPMI struct {
	IpmitoolInstalled bool      `json:"ipmitoolInstalled"`
	IpmitoolVersion   string    `json:"ipmitoolVersion"`
	IPSource          *string   `json:"ipSource,omitempty"`
	IPAddr            *string   `json:"ipAddr,omitempty"`
	DNSNames          *[]string `json:"dnsNames,omitempty"`
	Netmask           *string   `json:"netmask,omitempty"`
	HWAddr            *string   `json:"hwAddr,omitempty"`
	Gateway           *string   `json:"gateway,omitempty"`
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

	o, err := parse.Exec("ipmitool", []string{"lan", "print", "1"})
	if err == nil {
		for _, line := range strings.Split(o, "\n") {
			a := strings.SplitN(line, ":", 2)
			if len(a) < 2 {
				continue
			}

			k := strings.TrimSpace(a[0])
			v := strings.TrimSpace(a[1])

			switch k {
			case "IP Address Source":
				i.IPSource = &v
			case "IP Address":
				i.IPAddr = &v
			case "Subnet Mask":
				i.Netmask = &v
			case "MAC Address":
				i.HWAddr = &v
			case "Default Gateway IP":
				i.Gateway = &v
			}
		}
	}

	if i.IPAddr != nil {
		names, err := net.LookupAddr(*i.IPAddr)
		if err == nil {
			for i, v := range names {
				names[i] = strings.TrimRight(v, ".")
			}

			i.DNSNames = &names
		}
	}

	return i, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
