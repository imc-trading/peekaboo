package system

import (
	"time"
)

type System interface {
	GetData() Data
	GetCache() Cache
	GetDataIntf() interface{}
	GetCacheIntf() interface{}
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type system struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() System {
	return &system{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (s *system) GetData() Data {
	return *s.data
}

func (s *system) GetCache() Cache {
	return *s.cache
}

func (s *system) GetDataIntf() interface{} {
	return *s.data
}

func (s *system) GetCacheIntf() interface{} {
	return *s.cache
}

func (s *system) SetTimeout(timeout int) {
	s.cache.Timeout = timeout
}

func (s *system) Update() error {
	if s.cache.LastUpdated.IsZero() {
		if err := s.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := s.cache.LastUpdated.Add(time.Duration(s.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := s.ForceUpdate(); err != nil {
				return err
			}
		} else {
			s.cache.FromCache = true
		}
	}

	return nil
}
