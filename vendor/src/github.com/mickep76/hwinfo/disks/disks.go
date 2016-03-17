// +build linux

package disks

import (
	"path/filepath"
	"strconv"
	"time"

	"github.com/mickep76/hwinfo/common"
)

type Disks interface {
	GetData() Data
	GetCache() Cache
	GetDataIntf() interface{}
	GetCacheIntf() interface{}
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type disks struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data []dataItem

type dataItem struct {
	Device string `json:"device"`
	Name   string `json:"name"`
	SizeKB int    `json:"size_kb"`
	SizeGB int    `json:"size_gb"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Disks {
	return &disks{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (d *disks) GetData() Data {
	return *d.data
}

func (d *disks) GetCache() Cache {
	return *d.cache
}

func (d *disks) GetDataIntf() interface{} {
	return *d.data
}

func (d *disks) GetCacheIntf() interface{} {
	return *d.cache
}

func (d *disks) SetTimeout(timeout int) {
	d.cache.Timeout = timeout
}

func (d *disks) Update() error {
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

func (disks *disks) ForceUpdate() error {
	disks.cache.LastUpdated = time.Now()
	disks.cache.FromCache = false

	files, err := filepath.Glob("/sys/class/block/*")
	if err != nil {
		return err
	}

	for _, path := range files {
		o, err := common.LoadFiles([]string{
			filepath.Join(path, "dev"),
			filepath.Join(path, "size"),
		})
		if err != nil {
			return err
		}

		d := dataItem{}

		d.Name = filepath.Base(path)
		d.Device = o["dev"]

		d.SizeKB, err = strconv.Atoi(o["size"])
		if err != nil {
			return err
		}
		d.SizeKB = d.SizeKB * 512 / 1024
		d.SizeGB = d.SizeKB / 1024 / 1024

		*disks.data = append(*disks.data, d)
	}

	return nil
}
