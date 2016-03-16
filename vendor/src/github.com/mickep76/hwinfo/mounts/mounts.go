// +build linux

package mounts

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Mounts interface {
	GetData() data
	GetCache() cache
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type mounts struct {
	data  *data  `json:"data"`
	cache *cache `json:"cache"`
}

type data []dataItem

type dataItem struct {
	Source  string `json:"source"`
	Target  string `json:"target"`
	FSType  string `json:"fs_type"`
	Options string `json:"options"`
}

type cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Mounts {
	return &mounts{
		data: &data{},
		cache: &cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (m *mounts) GetData() data {
	return *m.data
}

func (m *mounts) GetCache() cache {
	return *m.cache
}

func (m *mounts) SetTimeout(timeout int) {
	m.cache.Timeout = timeout
}

func (m *mounts) Update() error {
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

func (mounts *mounts) ForceUpdate() error {
	mounts.cache.LastUpdated = time.Now()
	mounts.cache.FromCache = false

	fn := "/proc/mounts"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		m := dataItem{}

		m.Source = v[0]
		m.Target = v[1]
		m.FSType = v[2]
		m.Options = v[3]

		*mounts.data = append(*mounts.data, m)
	}

	return nil
}
