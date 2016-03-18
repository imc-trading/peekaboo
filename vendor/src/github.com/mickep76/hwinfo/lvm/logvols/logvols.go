// +build linux

package logvols

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type LogVols interface {
	GetData() Data
	GetCache() Cache
	GetDataIntf() interface{}
	GetCacheIntf() interface{}
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type logVols struct {
	data  *Data  `json:"data"`
	cache *Cache `json:"cache"`
}

type Data []DataItem

type DataItem struct {
	Name   string `json:"name"`
	VolGrp string `json:"vol_grp"`
	Attr   string `json:"attr"`
	SizeKB int    `json:"size_kb"`
	SizeGB int    `json:"size_gb"`
}

type Cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() LogVols {
	return &logVols{
		data: &Data{},
		cache: &Cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (l *logVols) GetData() Data {
	return *l.data
}

func (l *logVols) GetCache() Cache {
	return *l.cache
}

func (l *logVols) GetDataIntf() interface{} {
	return *l.data
}

func (l *logVols) GetCacheIntf() interface{} {
	return *l.cache
}

func (l *logVols) SetTimeout(timeout int) {
	l.cache.Timeout = timeout
}

func (l *logVols) Update() error {
	if l.cache.LastUpdated.IsZero() {
		if err := l.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := l.cache.LastUpdated.Add(time.Duration(l.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := l.ForceUpdate(); err != nil {
				return err
			}
		} else {
			l.cache.FromCache = true
		}
	}

	return nil
}

func (l *logVols) ForceUpdate() error {
	l.cache.LastUpdated = time.Now()
	l.cache.FromCache = false

	_, err := exec.LookPath("lvs")
	if err != nil {
		return errors.New("command doesn't exist: lvs")
	}

	o, err := exec.Command("lvs", "--units", "B").Output()
	if err != nil {
		return err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 1 || len(v) < 1 {
			continue
		}

		lv := DataItem{}

		lv.Name = v[0]
		lv.VolGrp = v[1]
		lv.Attr = v[2]

		lv.SizeKB, err = strconv.Atoi(strings.TrimRight(v[3], "B"))
		if err != nil {
			return err
		}
		lv.SizeKB = lv.SizeKB / 1024
		lv.SizeGB = lv.SizeKB / 1024 / 1024

		*l.data = append(*l.data, lv)
	}

	return nil
}
