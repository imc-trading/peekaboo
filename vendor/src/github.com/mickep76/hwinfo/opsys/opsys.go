package opsys

import (
	"time"
)

type OpSys interface {
	GetData() Data
	GetCache() Cache
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type opSys struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data struct {
	Kernel         string `json:"kernel"`
	KernelVersion  string `json:"kernel_version"`
	Product        string `json:"product"`
	ProductVersion string `json:"product_version"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() OpSys {
	return &opSys{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (o *opSys) GetData() Data {
	return *o.data
}

func (o *opSys) GetCache() Cache {
	return *o.cache
}

func (o *opSys) SetTimeout(timeout int) {
	o.cache.Timeout = timeout
}

func (o *opSys) Update() error {
	if o.cache.LastUpdated.IsZero() {
		if err := o.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := o.cache.LastUpdated.Add(time.Duration(o.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := o.ForceUpdate(); err != nil {
				return err
			}
		} else {
			o.cache.FromCache = true
		}
	}

	return nil
}
