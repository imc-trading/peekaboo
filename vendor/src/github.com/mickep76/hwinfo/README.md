
# hwinfo
    import "github.com/mickep76/hwinfo"







## type HWInfo
``` go
type HWInfo struct {
    Hostname string           `json:"hostname"`
    CPU      *cpu.CPU         `json:"cpu"`
    Memory   *memory.Memory   `json:"memory"`
    OpSys    *opsys.OpSys     `json:"opsys"`
    System   *system.System   `json:"system"`
    Network  *network.Network `json:"network"`
}
```
HWInfo information.









### func Get
``` go
func Get() (HWInfo, error)
```
Get information about a system.










- - -

# common
    import "github.com/mickep76/hwinfo/common"






## func ExecCmd
``` go
func ExecCmd(cmd string, args []string) (string, error)
```
ExecCmd returns output.


## func ExecCmdFields
``` go
func ExecCmdFields(cmd string, args []string, del string, fields []string) (map[string]string, error)
```
ExecCmdFields returns fields from output.


## func LoadFileFields
``` go
func LoadFileFields(fn string, del string, fields []string) (map[string]string, error)
```
LoadFileFields returns fields from file.


## func LoadFiles
``` go
func LoadFiles(files []string) (map[string]string, error)
```
LoadFiles returns field from multiple files.









- - -

# cpu
    import "github.com/mickep76/hwinfo/cpu"







## type CPU
``` go
type CPU struct {
    Model          string `json:"model"`
    Flags          string `json:"flags"`
    Logical        int    `json:"logical"`
    Physical       int    `json:"physical"`
    Sockets        int    `json:"sockets"`
    CoresPerSocket int    `json:"cores_per_socket"`
    ThreadsPerCore int    `json:"threads_per_core"`
}
```
CPU information.









### func Get
``` go
func Get() (CPU, error)
```
Get information about system CPU(s).










- - -

# disks
    import "github.com/mickep76/hwinfo/disks"







## type Disk
``` go
type Disk struct {
    Device string `json:"device"`
    Name   string `json:"name"`
    SizeGB int    `json:"size_gb"`
}
```
Disk information.

















- - -

# lvm
    import "github.com/mickep76/hwinfo/lvm"







## type LVM
``` go
type LVM struct {
    PhysVols *[]PhysVol `json:"phys_vols"`
    LogVols  *[]LogVol  `json:"log_vols"`
    VolGrps  *[]VolGrp  `json:"vol_grps"`
}
```










## type LogVol
``` go
type LogVol struct {
    Name   string `json:"name"`
    VolGrp string `json:"vol_grp"`
    Attr   string `json:"attr"`
    SizeGB int    `json:"size_gb"`
}
```










## type PhysVol
``` go
type PhysVol struct {
    Name   string `json:"name"`
    VolGrp string `json:"vol_group"`
    Format string `json:"format"`
    Attr   string `json:"attr"`
    SizeGB int    `json:"size_gb"`
    FreeGB int    `json:"free_gb"`
}
```










## type VolGrp
``` go
type VolGrp struct {
    Name   string `json:"name"`
    Attr   string `json:"attr"`
    SizeGB int    `json:"size_gb"`
    FreeGB int    `json:"free_gb"`
}
```
















- - -

# memory
    import "github.com/mickep76/hwinfo/memory"







## type Memory
``` go
type Memory struct {
    TotalGB int `json:"total_gb"`
}
```
Memory information.









### func Get
``` go
func Get() (Memory, error)
```
Get information about system memory.










- - -

# mounts
    import "github.com/mickep76/hwinfo/mounts"







## type Mount
``` go
type Mount struct {
    Source  string `json:"source"`
    Target  string `json:"target"`
    FSType  string `json:"fs_type"`
    Options string `json:"options"`
}
```
















- - -

# network
    import "github.com/mickep76/hwinfo/network"







## type Interface
``` go
type Interface struct {
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
```
Info structure for information about a systems network interfaces.











## type Network
``` go
type Network struct {
    Interfaces    []Interface `json:"interfaces"`
    OnloadVersion string      `json:"onload_version,omitempty"`
}
```
Info structure for information about a systems network.









### func Get
``` go
func Get() (Network, error)
```
GetInfo return information about a systems memory.










- - -

# opsys
    import "github.com/mickep76/hwinfo/opsys"







## type OpSys
``` go
type OpSys struct {
    Kernel         string `json:"kernel"`
    KernelVersion  string `json:"kernel_version"`
    Product        string `json:"product"`
    ProductVersion string `json:"product_version"`
}
```
OpSys information.









### func Get
``` go
func Get() (OpSys, error)
```
Get information about the operating system.










- - -

# pci
    import "github.com/mickep76/hwinfo/pci"






## func Get
``` go
func Get() ([]PCI, error)
```
Get information about system PCI slots.



## type PCI
``` go
type PCI struct {
    Slot      string `json:"slot"`
    ClassID   string `json:"class_id"`
    Class     string `json:"class"`
    VendorID  string `json:"vendor_id"`
    DeviceID  string `json:"device_id"`
    Vendor    string `json:"vendor"`
    Device    string `json:"device"`
    SVendorID string `json:"svendor_id"`
    SDeviceID string `json:"sdevice_id"`
    SName     string `json:"sname,omiempty"`
}
```
















- - -

# routes
    import "github.com/mickep76/hwinfo/routes"







## type Route
``` go
type Route struct {
    Destination string `json:"destination"`
    Gateway     string `json:"gateway"`
    Genmask     string `json:"genmask"`
    Flags       string `json:"flags"`
    MSS         int    `json:"mss"` // Maximum segment size
    Window      int    `json:"window"`
    IRTT        int    `json:"irtt"` // Initial round trip time
    Interface   string `json:"interface"`
}
```
Info structure for system routes.

















- - -

# run








- - -

# sysctl
    import "github.com/mickep76/hwinfo/sysctl"







## type Sysctl
``` go
type Sysctl struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
```
Sysctl structure for sysctl key/values.

















- - -

# system
    import "github.com/mickep76/hwinfo/system"







## type System
``` go
type System struct {
    Manufacturer   string `json:"manufacturer"`
    Product        string `json:"product"`
    ProductVersion string `json:"product_version"`
    SerialNumber   string `json:"serial_number"`
    BIOSVendor     string `json:"bios_vendor,omitempty"`
    BIOSDate       string `json:"bios_date,omitempty"`
    BIOSVersion    string `json:"bios_version,omitempty"`
    BootROMVersion string `json:"boot_rom_version,omitempty"`
    SMCVersion     string `json:"smc_version,omitempty"`
}
```
System information.









### func Get
``` go
func Get() (System, error)
```
Get information about a system.










- - -
