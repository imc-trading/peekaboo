package netinfo

// Info structure for information about a systems memory.
type Info struct {
	Name   string   `json:"name"`
	MTU    int      `json:"mtu"`
	IPAddr []string `json:"ipaddr"`
	HWAddr string   `json:"hwaddr"`
}
