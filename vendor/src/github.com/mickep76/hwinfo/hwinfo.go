package hwinfo

func (h *hwInfo) GetData() Data {
	return *h.data
}

func (h *hwInfo) GetCache() Cache {
	return *h.cache
}
