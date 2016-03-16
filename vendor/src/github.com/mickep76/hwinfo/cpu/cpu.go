package cpu

import (
	"time"
)

type CPU interface {
	GetData() Data
	GetCache() Cache
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type cpu struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data struct {
	Model          string `json:"model"`
	Flags          string `json:"flags"`
	Logical        int    `json:"logical"`
	Physical       int    `json:"physical"`
	Sockets        int    `json:"sockets"`
	CoresPerSocket int    `json:"cores_per_socket"`
	ThreadsPerCore int    `json:"threads_per_core"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() CPU {
	return &cpu{
		data: &Data{},
		cache: &Cache{
			Timeout: 12 * 60 * 60, // 12 hours
		},
	}
}

func (c *cpu) Update() error {
	if c.cache.LastUpdated.IsZero() {
		if err := c.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := c.cache.LastUpdated.Add(time.Duration(c.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := c.ForceUpdate(); err != nil {
				return err
			}
		} else {
			c.cache.FromCache = true
		}
	}

	return nil
}

func (c *cpu) SetTimeout(timeout int) {
	c.cache.Timeout = timeout
}

func (c *cpu) GetData() Data {
	return *c.data
}

func (c *cpu) GetCache() Cache {
	return *c.cache
}
