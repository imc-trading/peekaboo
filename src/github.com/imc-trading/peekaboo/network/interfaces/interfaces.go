package interfaces

import (
	"encoding/hex"
	"fmt"
	"net"
	"runtime"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

// Interfaces multiple network interfaes.
type Interfaces []Interface

// Interface single network interface.
type Interface struct {
	Name            string   `json:"name"`
	MTU             int      `json:"mtu"`
	IPAddr          []string `json:"ipAddr"`
	HWAddr          string   `json:"hwAddr"`
	Flags           []string `json:"flags"`
	Driver          *string  `json:"driver,omitempty"`
	DriverVersion   *string  `json:"driverVersion,omitempty"`
	FirmwareVersion *string  `json:"firmwareVersion,omitempty"`
	PCIBus          *string  `json:"pciBus,omitempty"`
	PCIBusURL       *string  `json:"pciBusURL,omitempty"`
	SwChassisID     *string  `json:"swChassisId,omitempty"`
	SwName          *string  `json:"swName,omitempty"`
	SwDescr         *string  `json:"swDescr,omitempty"`
	SwPortID        *string  `json:"swPortId,omitempty"`
	SwPortDescr     *string  `json:"swPortDescr,omitempty"`
	SwVLAN          *string  `json:"swVLan,omitempty"`
	TwinexCableSN   *string  `json:"twinexCableSN,omitempty"`
}

// Get network interfaces.
func Get() (Interfaces, error) {
	hasEthtool := false
	if err := parse.Exists("ethtool"); err == nil {
		hasEthtool = true
	}

	hasLldpctl := false
	if err := parse.Exists("lldpctl"); err == nil {
		hasLldpctl = true
	}

	hasSfctool := false
	if err := parse.Exists("sfctool"); err == nil {
		hasSfctool = true
	}

	rIntfs, err := net.Interfaces()
	if err != nil {
		return Interfaces{}, err
	}

	i := Interfaces{}
	for _, rIntf := range rIntfs {
		// Skip loopback interfaces
		if rIntf.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := rIntf.Addrs()
		if err != nil {
			return Interfaces{}, err
		}

		wIntf := Interface{
			Name:   rIntf.Name,
			HWAddr: rIntf.HardwareAddr.String(),
			MTU:    rIntf.MTU,
		}

		for _, addr := range addrs {
			wIntf.IPAddr = append(wIntf.IPAddr, addr.String())
		}

		if rIntf.Flags&net.FlagUp != 0 {
			wIntf.Flags = append(wIntf.Flags, "UP")
		}

		if rIntf.Flags&net.FlagBroadcast != 0 {
			wIntf.Flags = append(wIntf.Flags, "BROADCAST")
		}

		if rIntf.Flags&net.FlagPointToPoint != 0 {
			wIntf.Flags = append(wIntf.Flags, "POINTTOPOINT")
		}

		if rIntf.Flags&net.FlagMulticast != 0 {
			wIntf.Flags = append(wIntf.Flags, "MULTICAST")
		}

		if runtime.GOOS == "linux" && hasEthtool == true {
			m, err := parse.ExecRegexpMap("ethtool", []string{"-i", rIntf.Name}, ":", "\\S+:\\s\\S+")
			if err != nil {
				return Interfaces{}, err
			}

			s1 := m["driver"]
			wIntf.Driver = &s1
			s2 := m["version"]
			wIntf.DriverVersion = &s2
			s3 := m["firmware-version"]
			wIntf.FirmwareVersion = &s3
			if strings.HasPrefix(m["bus-info"], "0000:") {
				s4 := m["bus-info"]
				wIntf.PCIBus = &s4
				s5 := fmt.Sprintf("/pci/%v", m["bus-info"])
				wIntf.PCIBusURL = &s5
			}
		}

		var sn string
		if runtime.GOOS == "linux" && hasEthtool && hasSfctool && *wIntf.Driver == "sfc" {
			o, err := parse.Exec("sfctool", []string{"-m", rIntf.Name, "hex", "on", "offset", "0x0040", "length", "16"})
			if err != nil {
				return Interfaces{}, err
			}

			for _, line := range strings.Split(o, "\n") {
				arr := strings.SplitN(line, ":", 2)
				if len(arr) < 2 {
					continue
				}

				key := strings.TrimSpace(arr[0])
				val := strings.TrimSpace(arr[1])

				switch key {
				case "0x0040":
					b, err := hex.DecodeString(strings.Replace(val, " ", "", -1))
					if err != nil {
					}
					sn = strings.Replace(string(b), "\u0000", "", -1)
				}
			}
		}

		if runtime.GOOS == "linux" && hasEthtool && hasSfctool && *wIntf.Driver == "sfc" && sn != "" {
			o, err := parse.Exec("sfctool", []string{"-m", rIntf.Name, "hex", "on", "offset", "0x0078", "length", "1"})
			if err != nil {
				return Interfaces{}, err
			}

			for _, line := range strings.Split(o, "\n") {
				arr := strings.SplitN(line, ":", 2)
				if len(arr) < 2 {
					continue
				}

				key := strings.TrimSpace(arr[0])
				val := strings.TrimSpace(arr[1])

				switch key {
				case "0x0078":
					b, err := hex.DecodeString(strings.Replace(val, " ", "", -1))
					if err != nil {
					}
					s := strings.Replace(sn+":"+string(b), "\u0000", "", -1)
					wIntf.TwinexCableSN = &s
				}
			}
		}

		if runtime.GOOS == "linux" && hasLldpctl == true {
			o, err := parse.Exec("lldpctl", []string{rIntf.Name})
			if err != nil {
				return Interfaces{}, err
			}

			for _, line := range strings.Split(o, "\n") {
				arr := strings.SplitN(line, ":", 2)
				if len(arr) < 2 {
					continue
				}

				key := strings.TrimSpace(arr[0])
				val := strings.TrimSpace(arr[1])

				switch key {
				case "ChassisID":
					wIntf.SwChassisID = &val
				case "SysName":
					wIntf.SwName = &val
				case "SysDescr":
					wIntf.SwDescr = &val
				case "PortID":
					wIntf.SwPortID = &val
				case "PortDescr":
					wIntf.SwPortDescr = &val
				case "VLAN":
					wIntf.SwVLAN = &val
				}
			}
		}

		i = append(i, wIntf)
	}

	return i, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
