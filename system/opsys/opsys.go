package opsys

type OpSys struct {
	Kernel         string `json:"kernel"`
	KernelVersion  string `json:"kernelVersion"`
	Product        string `json:"product"`
	ProductVersion string `json:"productVersion"`
}

func GetInterface() (interface{}, error) {
	return Get()
}
