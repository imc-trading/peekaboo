package hwinfo

func (h *hwInfo) GetData() data {
	return *h.data
}

func (h *hwInfo) GetCache() cache {
	return *h.cache
}
