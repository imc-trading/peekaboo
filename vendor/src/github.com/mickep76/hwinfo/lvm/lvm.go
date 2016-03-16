// +build linux

package lvm

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type LVM interface {
	GetData() data
	GetCache() cache
	SetTimeout(int)
	Update() error
	ForceUpdate() error
}

type lvm struct {
	data  *data  `json:"data"`
	cache *cache `json:"cache"`
}

type physVols []physVol

type physVol struct {
	Name   string `json:"name"`
	VolGrp string `json:"vol_group"`
	Format string `json:"format"`
	Attr   string `json:"attr"`
	SizeKB int    `json:"size_kb"`
	SizeGB int    `json:"size_gb"`
	FreeKB int    `json:"free_kb"`
	FreeGB int    `json:"free_gb"`
}

type logVols []logVol

type logVol struct {
	Name   string `json:"name"`
	VolGrp string `json:"vol_grp"`
	Attr   string `json:"attr"`
	SizeKB int    `json:"size_kb"`
	SizeGB int    `json:"size_gb"`
}

type volGrps []volGrp

type volGrp struct {
	Name   string `json:"name"`
	Attr   string `json:"attr"`
	SizeKB int    `json:"size_kb"`
	SizeGB int    `json:"size_gb"`
	FreeKB int    `json:"free_kb"`
	FreeGB int    `json:"free_gb"`
}

type data struct {
	PhysVols physVols `json:"phys_vols"`
	LogVols  logVols  `json:"log_vols"`
	VolGrps  volGrps  `json:"vol_grps"`
}

type cache struct {
	LastUpdated time.Time `json:"last_updated"`
	Timeout     int       `json:"timeout_sec"`
	FromCache   bool      `json:"from_cache"`
}

func New() LVM {
	return &lvm{
		data: &data{
			PhysVols: physVols{},
			LogVols:  logVols{},
			VolGrps:  volGrps{},
		},
		cache: &cache{
			Timeout: 5 * 60, // 5 minutes
		},
	}
}

func (l *lvm) GetData() data {
	return *l.data
}

func (l *lvm) GetCache() cache {
	return *l.cache
}

func (l *lvm) SetTimeout(timeout int) {
	l.cache.Timeout = timeout
}

func (l *lvm) Update() error {
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

func (l *lvm) ForceUpdate() error {
	if err := l.getPhysVols(); err != nil {
		return err
	}

	if err := l.getLogVols(); err != nil {
		return err
	}

	if err := l.getVolGrps(); err != nil {
		return err
	}

	return nil
}

func (l *lvm) getPhysVols() error {
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

		pv := physVol{}

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

		l.data.PhysVols = append(l.data.PhysVols, pv)
	}

	return nil
}

func (l *lvm) getLogVols() error {
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

		lv := logVol{}

		lv.Name = v[0]
		lv.VolGrp = v[1]
		lv.Attr = v[2]

		lv.SizeKB, err = strconv.Atoi(strings.TrimRight(v[3], "B"))
		if err != nil {
			return err
		}
		lv.SizeKB = lv.SizeKB / 1024
		lv.SizeGB = lv.SizeKB / 1024 / 1024

		l.data.LogVols = append(l.data.LogVols, lv)
	}

	return nil
}

func (l *lvm) getVolGrps() error {
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

		vg := volGrp{}

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

		l.data.VolGrps = append(l.data.VolGrps, vg)
	}

	return nil
}
