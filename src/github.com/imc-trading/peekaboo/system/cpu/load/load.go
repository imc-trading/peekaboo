package load

// Load structure.
type Load struct {
	Avg1    float32 `json:"loadAvg1"`
	Avg5    float32 `json:"loadAvg5"`
	Avg10   float32 `json:"loadAvg10"`
	Running int     `json:"processesRunning"`
	Total   int     `json:"processesTotal"`
	LastPid int     `json:"processesLastPid"`
}

func GetInterface() (interface{}, error) {
	return Get()
}
