package diskinfo

// Info structure for information about a systems memory.
type Info struct {
	Disks []Disk `json:"disk"`
}

type Disk struct {
	Device string `json:"device"`
	Name   string `json:"name"`
	//	Major  int    `json:"major"`
	//	Minor  int    `json:"minor"`
	//	Blocks int    `json:"blocks"`
	SizeGB int `json:"size_gb"`
}
