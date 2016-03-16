package interfaces

import (
	"fmt"
	"net"
	"runtime"
	"strings"
	"time"

	"github.com/mickep76/hwinfo/common"
)

type Interfaces interface {
	GetData() Data
	GetCache() Cache
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type interfaces struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data []DataItem

type DataItem struct {
	Name            string   `json:"name"`
	MTU             int      `json:"mtu"`
	IPAddr          []string `json:"ipaddr"`
	HWAddr          string   `json:"hwaddr"`
	Flags           []string `json:"flags"`
	Driver          string   `json:"driver,omitempty"`
	DriverVersion   string   `json:"driver_version,omitempty"`
	FirmwareVersion string   `json:"firmware_version,omitempty"`
	PCIBus          string   `json:"pci_bus,omitempty"`
	PCIBusURL       string   `json:"pci_bus_url,omitempty"`
	SwChassisID     string   `json:"sw_chassis_id"`
	SwName          string   `json:"sw_name"`
	SwDescr         string   `json:"sw_descr"`
	SwPortID        string   `json:"sw_port_id"`
	SwPortDescr     string   `json:"sw_port_descr"`
	SwVLAN          string   `json:"sw_vlan"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Interfaces {
	return &interfaces{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (i *interfaces) GetData() Data {
	return *i.data
}

func (i *interfaces) GetCache() Cache {
	return *i.cache
}

func (i *interfaces) SetTimeout(timeout int) {
	i.cache.Timeout = timeout
}

func (i *interfaces) Update() error {
	if i.cache.LastUpdated.IsZero() {
		if err := i.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := i.cache.LastUpdated.Add(time.Duration(i.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := i.ForceUpdate(); err != nil {
				return err
			}
		} else {
			i.cache.FromCache = true
		}
	}

	return nil
}

func (i *interfaces) ForceUpdate() error {
	i.cache.LastUpdated = time.Now()
	i.cache.FromCache = false

	rIntfs, err := net.Interfaces()
	if err != nil {
		return err
	}

	for _, rIntf := range rIntfs {
		// Skip loopback interfaces
		if rIntf.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := rIntf.Addrs()
		if err != nil {
			return err
		}

		sIntf := DataItem{
			Name:   rIntf.Name,
			HWAddr: rIntf.HardwareAddr.String(),
			MTU:    rIntf.MTU,
		}

		for _, addr := range addrs {
			sIntf.IPAddr = append(sIntf.IPAddr, addr.String())
		}

		if rIntf.Flags&net.FlagUp != 0 {
			sIntf.Flags = append(sIntf.Flags, "up")
		}
		if rIntf.Flags&net.FlagBroadcast != 0 {
			sIntf.Flags = append(sIntf.Flags, "broadcast")
		}
		if rIntf.Flags&net.FlagPointToPoint != 0 {
			sIntf.Flags = append(sIntf.Flags, "pointtopoint")
		}
		if rIntf.Flags&net.FlagMulticast != 0 {
			sIntf.Flags = append(sIntf.Flags, "multicast")
		}

		switch runtime.GOOS {
		case "linux":
			o, err := common.ExecCmdFields("/usr/sbin/ethtool", []string{"-i", rIntf.Name}, ":", []string{
				"driver",
				"version",
				"firmware-version",
				"bus-info",
			})
			if err != nil {
				return err
			}

			sIntf.Driver = o["driver"]
			sIntf.DriverVersion = o["version"]
			sIntf.FirmwareVersion = o["firmware-version"]
			if strings.HasPrefix(o["bus-info"], "0000:") {
				sIntf.PCIBus = o["bus-info"]
				sIntf.PCIBusURL = fmt.Sprintf("/pci/%v", o["bus-info"])
			}

			o2, err := common.ExecCmdFields("/usr/sbin/lldpctl", []string{rIntf.Name}, ":", []string{
				"ChassisID",
				"SysName",
				"SysDescr",
				"PortID",
				"PortDescr",
				"VLAN",
			})
			if err != nil {
				return err
			}

			sIntf.SwChassisID = o2["ChassisID"]
			sIntf.SwName = o2["SysName"]
			sIntf.SwDescr = o2["SysDescr"]
			sIntf.SwPortID = o2["PortID"]
			sIntf.SwPortDescr = o2["PortDescr"]
			sIntf.SwVLAN = o2["VLAN"]
		}

		*i.data = append(*i.data, sIntf)
	}

	return nil
}
