// +build linux

package volgrps

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type VolGrps interface {
	GetData() Data
	GetCache() Cache
	GetDataIntf() interface{}
	GetCacheIntf() interface{}
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type volGrps struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data []DataItem

type DataItem struct {
	Name   string `json:"name"`
	Attr   string `json:"attr"`
	SizeKB int    `json:"size_kb"`
	SizeGB int    `json:"size_gb"`
	FreeKB int    `json:"free_kb"`
	FreeGB int    `json:"free_gb"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() VolGrps {
	return &volGrps{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (v *volGrps) GetData() Data {
	return *v.data
}

func (v *volGrps) GetCache() Cache {
	return *v.cache
}

func (v *volGrps) GetDataIntf() interface{} {
	return *v.data
}

func (v *volGrps) GetCacheIntf() interface{} {
	return *v.cache
}

func (v *volGrps) SetTimeout(timeout int) {
	v.cache.Timeout = timeout
}

func (v *volGrps) Update() error {
	if v.cache.LastUpdated.IsZero() {
		if err := v.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := v.cache.LastUpdated.Add(time.Duration(v.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := v.ForceUpdate(); err != nil {
				return err
			}
		} else {
			v.cache.FromCache = true
		}
	}

	return nil
}

func (vgs *volGrps) ForceUpdate() error {
	vgs.cache.LastUpdated = time.Now()
	vgs.cache.FromCache = false

	_, err := exec.LookPath("vgs")
	if err != nil {
		return errors.New("command doesn't exist: vgs")
	}

	o, err := exec.Command("vgs", "--units", "B").Output()
	if err != nil {
		return err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		vg := DataItem{}

		vg.Name = v[0]
		vg.Attr = v[4]

		vg.SizeKB, err = strconv.Atoi(strings.TrimRight(v[5], "B"))
		if err != nil {
			return err
		}
		vg.SizeKB = vg.SizeKB / 1024
		vg.SizeGB = vg.SizeKB / 1024 / 1024

		vg.FreeKB, err = strconv.Atoi(strings.TrimRight(v[6], "B"))
		if err != nil {
			return err
		}
		vg.FreeKB = vg.FreeKB / 1024
		vg.FreeGB = vg.FreeGB / 1024 / 1024

		*vgs.data = append(*vgs.data, vg)
	}

	return nil
}
