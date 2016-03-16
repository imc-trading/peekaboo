package memory

import (
	"time"
)

type Memory interface {
	GetData() Data
	GetCache() Cache
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type memory struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Memory {
	return &memory{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (m *memory) GetData() Data {
	return *m.data
}

func (m *memory) GetCache() Cache {
	return *m.cache
}

func (m *memory) SetTimeout(timeout int) {
	m.cache.Timeout = timeout
}

func (m *memory) Update() error {
	if m.cache.LastUpdated.IsZero() {
		if err := m.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := m.cache.LastUpdated.Add(time.Duration(m.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := m.ForceUpdate(); err != nil {
				return err
			}
		} else {
			m.cache.FromCache = true
		}
	}

	return nil
}
