package interfaces

import (
	"net"
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
	Driver          string   `json:"driver,omitempty"`
	DriverVersion   string   `json:"driverVersion,omitempty"`
	FirmwareVersion string   `json:"firmwareVersion,omitempty"`
	PCIBus          string   `json:"pciBus,omitempty"`
	PCIBusURL       string   `json:"pciBusURL,omitempty"`
	SwChassisID     string   `json:"swChassisId"`
	SwName          string   `json:"swName"`
	SwDescr         string   `json:"swDescr"`
	SwPortID        string   `json:"swPortId"`
	SwPortDescr     string   `json:"swPortDescr"`
	SwVLAN          string   `json:"swVLan"`
}

// Get network interfaces.
func Get() (Interfaces, error) {
	i := Interfaces{}

	rIntfs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, rIntf := range rIntfs {
		// Skip loopback interfaces
		if rIntf.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := rIntf.Addrs()
		if err != nil {
			return nil, err
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

		/*
		   		switch runtime.GOOS {
		   		case "linux":
		   		    m, err := parse.ExecRegexpMap("/usr/sbin/ethtool", []string{"-i", rIntf.Name}, ":", "\\S+:\\s\\S+")
		   		    if err != nil {
		   				return err
		   		    }

		   			wIntf.Driver = m["driver"]
		   			wIntf.DriverVersion = m["version"]
		   			wIntf.FirmwareVersion = m["firmware-version"]
		   			if strings.HasPrefix(m["bus-info"], "0000:") {
		   				wIntf.PCIBus = m["bus-info"]
		   				wIntf.PCIBusURL = fmt.Sprintf("/pci/%v", o["bus-info"])
		   			}

		   			m, err := parse.ExecRegexpMap("/usr/sbin/lldpctl", []string{rIntf.Name}, ":", "\\S+:\\s\\S+")
		       if err != nil {
		   				return err
		   			}

		   			wIntf.SwChassisID = o2["ChassisID"]
		   			wIntf.SwName = o2["SysName"]
		   			wIntf.SwDescr = o2["SysDescr"]
		   			wIntf.SwPortID = o2["PortID"]
		   			wIntf.SwPortDescr = o2["PortDescr"]
		   			wIntf.SwVLAN = o2["VLAN"]
		*/

		i = append(i, wIntf)
	}

	return i, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
