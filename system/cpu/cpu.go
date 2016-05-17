package cpu

// CPU structure.
type CPU struct {
	Model          string `json:"model"`
	Flags          string `json:"flags"`
	Logical        int    `json:"logical"`
	Physical       int    `json:"physical"`
	Sockets        int    `json:"sockets"`
	CoresPerSocket int    `json:"coresPerSocket"`
	ThreadsPerCore int    `json:"threadsPerCore"`
}

func GetInterface() (interface{}, error) {
	return Get()
}
