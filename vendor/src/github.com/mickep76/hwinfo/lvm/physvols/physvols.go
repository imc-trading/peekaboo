// +build linux

package physvols

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type PhysVols interface {
	GetData() Data
	GetCache() Cache
	GetDataIntf() interface{}
	GetCacheIntf() interface{}
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type physVols struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data []DataItem

type DataItem struct {
	Name   string `json:"name"`
	VolGrp string `json:"vol_group"`
	Format string `json:"format"`
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

func New() PhysVols {
	return &physVols{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (p *physVols) GetData() Data {
	return *p.data
}

func (p *physVols) GetCache() Cache {
	return *p.cache
}

func (p *physVols) GetDataIntf() interface{} {
	return *p.data
}

func (p *physVols) GetCacheIntf() interface{} {
	return *p.cache
}

func (p *physVols) SetTimeout(timeout int) {
	p.cache.Timeout = timeout
}

func (p *physVols) Update() error {
	if p.cache.LastUpdated.IsZero() {
		if err := p.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := p.cache.LastUpdated.Add(time.Duration(p.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := p.ForceUpdate(); err != nil {
				return err
			}
		} else {
			p.cache.FromCache = true
		}
	}

	return nil
}

func (p *physVols) ForceUpdate() error {
	p.cache.LastUpdated = time.Now()
	p.cache.FromCache = false

	_, err := exec.LookPath("pvs")
	if err != nil {
		return errors.New("command doesn't exist: pvs")
	}

	o, err := exec.Command("pvs", "--units", "B").Output()
	if err != nil {
		return err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		pv := DataItem{}

		pv.Name = v[0]
		pv.VolGrp = v[1]
		pv.Format = v[2]
		pv.Attr = v[3]

		pv.SizeKB, err = strconv.Atoi(strings.TrimRight(v[4], "B"))
		if err != nil {
			return err
		}
		pv.SizeKB = pv.SizeKB / 1024
		pv.SizeGB = pv.SizeKB / 1024 / 1024

		pv.FreeKB, err = strconv.Atoi(strings.TrimRight(v[5], "B"))
		if err != nil {
			return err
		}
		pv.FreeKB = pv.FreeKB / 1024
		pv.FreeGB = pv.FreeKB / 1024 / 1024

		*p.data = append(*p.data, pv)
	}

	return nil
}
