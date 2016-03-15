package memory

// Memory information.
type Memory struct {
	TotalGB            int `json:"total_gb"`
	FreeGB             int `json:"free_gb"`
	AvailableGB        int `json:"available_gb"`
	CachedGB           int `json:"cached_gb"`
	CommittedActSizeGB int `json:"committed_act_size_gb"`
	HugePagesTot       int `json:"huge_pages_tot"`
	HugePagesFree      int `json:"huge_pages_free"`
	HugePagesRsvd      int `json:"huge_pages_rsvd"`
	HugePagesSurp      int `json:"huge_pages_surp"`
	HugePageSizeKB     int `json:"huge_pages_size_kb"`
}
