package osinfo

// Info structure for information about the operating system.
type Info struct {
	Kernel         string `json:"kernel"`
	KernelVersion  string `json:"kernel_version"`
	Product        string `json:"product"`
	ProductVersion string `json:"product_version"`
}
