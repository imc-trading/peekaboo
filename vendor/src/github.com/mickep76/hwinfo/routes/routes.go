// +build linux

package routes

import (
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Routes interface {
	GetData() data
	GetCache() cache
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type routes struct {
	data  *data  `json:"data"`
	cache *cache `json:"cache"`
}

type data []dataItem

type dataItem struct {
	Destination string `json:"destination"`
	Gateway     string `json:"gateway"`
	Genmask     string `json:"genmask"`
	Flags       string `json:"flags"`
	MSS         int    `json:"mss"` // Maximum segment size
	Window      int    `json:"window"`
	IRTT        int    `json:"irtt"` // Initial round trip time
	Interface   string `json:"interface"`
}

type cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() Routes {
	return &routes{
		data: &data{},
		cache: &cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (r *routes) GetData() data {
	return *r.data
}

func (r *routes) GetCache() cache {
	return *r.cache
}

func (r *routes) SetTimeout(timeout int) {
	r.cache.Timeout = timeout
}

func (r *routes) Update() error {
	if r.cache.LastUpdated.IsZero() {
		if err := r.ForceUpdate(); err != nil {
			return err
		}
	} else {
		expire := r.cache.LastUpdated.Add(time.Duration(r.cache.Timeout) * time.Second)
		if expire.Before(time.Now()) {
			if err := r.ForceUpdate(); err != nil {
				return err
			}
		} else {
			r.cache.FromCache = true
		}
	}

	return nil
}

func (routes *routes) ForceUpdate() error {
	routes.cache.LastUpdated = time.Now()
	routes.cache.FromCache = false

	o, err := exec.Command("netstat", "-rn").Output()
	if err != nil {
		return err
	}

	for c, line := range strings.Split(string(o), "\n") {
		v := strings.Fields(line)
		if c < 2 || len(v) < 8 {
			continue
		}

		r := dataItem{}

		r.Destination = v[0]
		r.Gateway = v[1]
		r.Genmask = v[2]
		r.Flags = v[3]

		r.MSS, err = strconv.Atoi(v[4])
		if err != nil {
			return err
		}

		r.Window, err = strconv.Atoi(v[5])
		if err != nil {
			return err
		}

		r.IRTT, err = strconv.Atoi(v[6])
		if err != nil {
			return err
		}

		r.Interface = v[7]

		*routes.data = append(*routes.data, r)
	}

	return nil
}
