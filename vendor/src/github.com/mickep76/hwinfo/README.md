
# hwinfo
    import "github.com/mickep76/hwinfo"







## type Info
``` go
type Info struct {
    Hostname string         `json:"hostname"`
    CPU      *cpuinfo.Info  `json:"cpu"`
    Memory   *meminfo.Info  `json:"memory"`
    OS       *osinfo.Info   `json:"os"`
    System   *sysinfo.Info  `json:"system"`
    Network  *netinfo.Info  `json:"network"`
    PCI      *pciinfo.Info  `json:"pci,omitempty"`
    Disk     *diskinfo.Info `json:"disk"`
}
```
Info structure for information a system.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a system.










- - -

# cpuinfo
    import "github.com/mickep76/hwinfo/cpuinfo"







## type Info
``` go
type Info struct {
    Model          string `json:"model"`
    Flags          string `json:"flags"`
    Logical        int    `json:"logical"`
    Physical       int    `json:"physical"`
    Sockets        int    `json:"sockets"`
    CoresPerSocket int    `json:"cores_per_socket"`
    ThreadsPerCore int    `json:"threads_per_core"`
}
```
Info structure for information about a systems CPU(s).









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a systems CPU(s).










- - -

# meminfo
    import "github.com/mickep76/hwinfo/meminfo"







## type Info
``` go
type Info struct {
    TotalKB int `json:"total_kb"`
}
```
Info structure for information about a systems memory.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a systems memory.










- - -

# osinfo
    import "github.com/mickep76/hwinfo/osinfo"







## type Info
``` go
type Info struct {
    Kernel         string `json:"kernel"`
    KernelVersion  string `json:"kernel_version"`
    Product        string `json:"product"`
    ProductVersion string `json:"product_version"`
}
```
Info structure for information about the operating system.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about the operating system.










- - -

# sysinfo
    import "github.com/mickep76/hwinfo/sysinfo"







## type Info
``` go
type Info struct {
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
Info structure for information about a system.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about a systems memory.










- - -

# pciinfo
    import "github.com/mickep76/hwinfo/pciinfo"







## type Info
``` go
type Info struct {
    PCI []PCI `json:"pci"`
}
```
Info structure for information about a systems memory.









### func GetInfo
``` go
func GetInfo() (Info, error)
```
GetInfo return information about PCI devices.




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

# diskinfo
    import "github.com/mickep76/hwinfo/diskinfo"







## type Disk
``` go
type Disk struct {
    Device string `json:"device"`
    Name   string `json:"name"`
    //	Major  int    `json:"major"`
    //	Minor  int    `json:"minor"`
    //	Blocks int    `json:"blocks"`
    SizeGB int `json:"size_gb"`
}
```










## type Info
``` go
type Info struct {
    Disks []Disk `json:"disk"`
}
```
Info structure for information about a systems memory.









### func GetInfo
``` go
func GetInfo() (Info, error)
```









- - -
