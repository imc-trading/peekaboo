// +build linux

package sysctl

import (
	"os/exec"
	"strings"
	"time"
)

type Sysctl interface {
	GetData() Data
	GetCache() Cache
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type sysctl struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data []dataItem

type dataItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Sysctl {
	return &sysctl{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (s *sysctl) GetData() Data {
	return *s.data
}

func (s *sysctl) GetCache() Cache {
	return *s.cache
}

func (s *sysctl) SetTimeout(timeout int) {
	s.cache.Timeout = timeout
}

func (s *sysctl) Update() error {
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

func (sysctl *sysctl) ForceUpdate() error {
	sysctl.cache.LastUpdated = time.Now()
	sysctl.cache.FromCache = false

	o, err := exec.Command("sysctl", "-a").Output()
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.Fields(line)
		if len(vals) < 3 {
			continue
		}

		s := dataItem{}

		s.Key = vals[0]
		s.Value = vals[2]

		*sysctl.data = append(*sysctl.data, s)
	}

	return nil
}
