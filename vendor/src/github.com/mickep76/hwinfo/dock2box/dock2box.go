// +build linux

package dock2box

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Dock2Box interface {
	GetData() Data
	GetCache() Cache
	GetDataIntf() interface{}
	GetCacheIntf() interface{}
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type dock2box struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data struct {
	FirstBoot string `json:"firstboot"`
	CFlags    string `json:"cflags_march_native"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Dock2Box {
	return &dock2box{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (d *dock2box) GetData() Data {
	return *d.data
}

func (d *dock2box) GetCache() Cache {
	return *d.cache
}

func (d *dock2box) GetDataIntf() interface{} {
	return *d.data
}

func (d *dock2box) GetCacheIntf() interface{} {
	return *d.cache
}

func (d *dock2box) SetTimeout(timeout int) {
	d.cache.Timeout = timeout
}

func (d *dock2box) Update() error {
	if d.cache.LastUpdated.IsZero() {
		if err := d.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := d.cache.LastUpdated.Add(time.Duration(d.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := d.ForceUpdate(); err != nil {
				return err
			}
		} else {
			d.cache.FromCache = true
		}
	}

	return nil
}

func (d *dock2box) ForceUpdate() error {
	d.cache.LastUpdated = time.Now()
	d.cache.FromCache = false

	file := "/etc/dock2box/firstboot.json"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return fmt.Errorf("file doesn't exist: %s", file)
	}

	o, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(o, d.data); err != nil {
		return err
	}

	return nil
}
