// +build darwin

package system

import (
	"os"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
	"github.com/imc-trading/peekaboo/version"
)

type System struct {
	Hostname        string `json:"hostname"`
	ShortHostname   string `json:"shortHostname"`
	PeekabooVersion string `json:"peekabooVersion"`
	Manufacturer    string `json:"manufacturer"`
	Product         string `json:"product"`
	ProductVersion  string `json:"productVersion"`
	SerialNumber    string `json:"serialNumber"`
	BootROMVersion  string `json:"bootROMVersion"`
	SMCVersion      string `json:"smcVersion"`
}

func Get() (System, error) {
	s := System{}

	m, err := parse.ExecRegexpMap("/usr/sbin/system_profiler", []string{"SPHardwareDataType"}, ":", "\\S+:\\s\\S+")
	if err != nil {
		return System{}, err
	}

	host, err := os.Hostname()
	if err != nil {
		return System{}, err
	}
	s.Hostname = host
	s.ShortHostname = strings.Split(host, ".")[0]

	s.PeekabooVersion = version.Version

	s.Manufacturer = "Apple Inc."
	s.Product = m["Model Name"]
	s.ProductVersion = m["Model Identifier"]
	s.SerialNumber = m["Serial Number (system)"]
	s.BootROMVersion = m["Boot ROM Version"]
	s.SMCVersion = m["SMC Version (system)"]

	return s, nil
}
