package interfaces

import (
	"encoding/hex"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

// Interfaces multiple network interfaes.
type Interfaces []Interface

// Interface single network interface.
type Interface struct {
	Name              string   `json:"name"`
	MTU               int      `json:"mtu"`
	IPAddr            []string `json:"ipAddr"`
	PermanentHWAddr   string   `json:"permanentHwAddr,omitempty"`
	HWAddr            string   `json:"hwAddr"`
	Flags             []string `json:"flags"`
	Driver            *string  `json:"driver,omitempty"`
	DriverVersion     *string  `json:"driverVersion,omitempty"`
	FirmwareVersion   *string  `json:"firmwareVersion,omitempty"`
	PCIBus            *string  `json:"pciBus,omitempty"`
	PCIBusURL         *string  `json:"pciBusURL,omitempty"`
	SpeedMbs          *int     `json:"speedMbs,omitempty"`
	Duplex            *string  `json:"duplex,omitempty"`
	LinkDetected      *bool    `json:"linkDetected,omitempty"`
	SwChassisID       *string  `json:"swChassisId,omitempty"`
	SwName            *string  `json:"swName,omitempty"`
	SwDescr           *string  `json:"swDescr,omitempty"`
	SwPortID          *string  `json:"swPortId,omitempty"`
	SwPortDescr       *string  `json:"swPortDescr,omitempty"`
	SwVLAN            *string  `json:"swVLan,omitempty"`
	TransceiverSN     *string  `json:"transceiverSN,omitempty"`
	TransceiverSA     *int     `json:"transceiverSA,omitempty"`
	TransceiverVN     *string  `json:"transceiverVN,omitempty"`
	TransceiverPN     *string  `json:"transceiverPN,omitempty"`
	TransceiverEeprom *string  `json:"transceiverEeprom,omitempty"`
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

			// Do nothing on error
			if err == nil {
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
		}

		if runtime.GOOS == "linux" && hasEthtool == true {
			m, err := parse.ExecRegexpMap("ethtool", []string{rIntf.Name}, ":", "\\S+:\\s\\S+")

			// Do nothing on error
			if err == nil {
				i, err := strconv.Atoi(strings.TrimRight(m["Speed"], "Mb/s"))
				if err == nil {
					wIntf.SpeedMbs = &i
				}

				if wIntf.SpeedMbs != nil {
					s := m["Duplex"]
					wIntf.Duplex = &s
				}

				switch m["Link detected"] {
				case "yes":
					b := true
					wIntf.LinkDetected = &b
				case "no":
					b := false
					wIntf.LinkDetected = &b
				}
			}
		}

		if runtime.GOOS == "linux" && hasEthtool == true {
			m, err := parse.ExecRegexpMap("ethtool", []string{"-P", rIntf.Name}, ":[ ]", "\\S+:\\s([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})")

			// Do nothing on error
			if err == nil {
				bia := m["Permanent address"]
				wIntf.PermanentHWAddr = bia
			}
		}

		if runtime.GOOS == "linux" && hasEthtool && wIntf.Driver != nil && *wIntf.Driver != "sfc" {
			o, err := parse.Exec("ethtool", []string{"-m", rIntf.Name, "raw", "on"})

			// Do nothing on error, doesn't support getting Eeprom info
			if err == nil {
				s := hex.EncodeToString([]byte(o))

				var b []byte
				var eeprom string = ""

				switch s[:2] {
				case "03":   // SFP+
					eeprom = s[0:]
				case "0d":   // QSFP+
					eeprom = s[256:768]
				}
				
				length := len(eeprom)
				if length >= 512 {
					ep := eeprom
					wIntf.TransceiverEeprom = &ep

					// Serial Number
					b, err = hex.DecodeString(eeprom[136:168])
					if err == nil {
						sn := strings.Trim(string(b), " ")
						wIntf.TransceiverSN = &sn
					}
					
					// Vendor Name
					b, err = hex.DecodeString(eeprom[40:72])
					if err == nil {
						vn := strings.Trim(string(b), " ")
						wIntf.TransceiverVN = &vn
					}
					
					// Product Name
					b, err = hex.DecodeString(eeprom[80:112])
					if err == nil {
						pn := strings.Trim(string(b), " ")
						wIntf.TransceiverPN = &pn
					}
					
					// Arista specific QSFP-4xSFP sub-assembly
					if *wIntf.TransceiverVN == "Arista Networks" && strings.HasPrefix(*wIntf.TransceiverPN, "CAB-Q-S-") {
						b, err = hex.DecodeString(eeprom[240:242])
						if err == nil {
							sa := int(b[0])
							wIntf.TransceiverSA = &sa
						}
					}
				}
			}
		}

		if runtime.GOOS == "linux" && hasSfctool && wIntf.Driver != nil && *wIntf.Driver == "sfc" {
			o, err := parse.Exec("sfctool", []string{"-m", rIntf.Name, "raw", "on"})

			// Do nothing on error, doesn't support getting Eeprom info
			if err == nil {
				s := hex.EncodeToString([]byte(o))

				var b []byte
				var eeprom string = ""

				switch s[:2] {
				case "03":   // SFP+
					eeprom = s[0:512]
				case "0d":   // QSFP+
					eeprom = s[256:768]
				}
				
				length := len(eeprom)
				if length == 512 {
					ep := eeprom
					wIntf.TransceiverEeprom = &ep

					// Serial Number
					b, err = hex.DecodeString(eeprom[136:168])
					if err == nil {
						sn := strings.Trim(string(b), " ")
						wIntf.TransceiverSN = &sn
					}
					
					// Vendor Name
					b, err = hex.DecodeString(eeprom[40:72])
					if err == nil {
						vn := strings.Trim(string(b), " ")
						wIntf.TransceiverVN = &vn
					}
					
					// Product Name
					b, err = hex.DecodeString(eeprom[80:112])
					if err == nil {
						pn := strings.Trim(string(b), " ")
						wIntf.TransceiverPN = &pn
					}
					
					// Arista specific QSFP-4xSFP sub-assembly
					if *wIntf.TransceiverVN == "Arista Networks" && strings.HasPrefix(*wIntf.TransceiverPN, "CAB-Q-S-") {
						b, err = hex.DecodeString(eeprom[240:242])
						if err == nil {
							sa := int(b[0])
							wIntf.TransceiverSA = &sa
						}
					}
				}
			}
		}

		if runtime.GOOS == "linux" && hasLldpctl == true {
			o, err := parse.Exec("lldpctl", []string{rIntf.Name})

			// Do nothing on error
			if err == nil {
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
		}

		i = append(i, wIntf)
	}

	return i, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
