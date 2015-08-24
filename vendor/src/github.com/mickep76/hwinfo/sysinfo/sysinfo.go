package sysinfo

// Info structure for information about a system.
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
