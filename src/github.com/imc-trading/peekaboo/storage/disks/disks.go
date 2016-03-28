// +build linux

package disks

import (
	"path/filepath"
	"strconv"

	"github.com/imc-trading/peekaboo/parse"
)

type Disks []Disk

type Disk struct {
	Device string  `json:"device"`
	Name   string  `json:"name"`
	SizeKB int     `json:"sizeKB"`
	SizeGB float32 `json:"sizeGB"`
}

func Get() (Disks, error) {
	disks := Disks{}

	files, err := filepath.Glob("/sys/class/block/*")
	if err != nil {
		return Disks{}, err
	}

	for _, path := range files {
		o, err := parse.LoadFiles([]string{
			filepath.Join(path, "dev"),
			filepath.Join(path, "size"),
		})
		if err != nil {
			return Disks{}, err
		}

		d := Disk{}

		d.Name = filepath.Base(path)
		d.Device = o["dev"]

		d.SizeKB, err = strconv.Atoi(o["size"])
		if err != nil {
			return Disks{}, err
		}
		d.SizeKB = d.SizeKB * 512 / 1024
		d.SizeGB = float32(d.SizeKB) / 1024 / 1024

		disks = append(disks, d)
	}

	return disks, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
