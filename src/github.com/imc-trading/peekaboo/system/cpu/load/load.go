package load

// Load structure.
type Load struct {
	Avg1    float32 `json:"loadAvg1"`
	Avg5    float32 `json:"loadAvg5"`
	Avg15   float32 `json:"loadAvg15"`
	Running int     `json:"processesRunning,omitempty"`
	Total   int     `json:"processesTotal,omitempty"`
	LastPid int     `json:"processesLastPid,omitempty"`
}

func GetInterface() (interface{}, error) {
	return Get()
}
