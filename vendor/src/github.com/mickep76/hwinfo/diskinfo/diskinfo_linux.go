// +build linux

package diskinfo

import (
	"github.com/mickep76/hwinfo/common"
	"path/filepath"
	"strconv"
)

func GetInfo() (Info, error) {
	i := Info{}

	files, err := filepath.Glob("/sys/class/block/*")
	if err != nil {
		return Info{}, err
	}

	for _, path := range files {
		o, err := common.LoadFiles([]string{
			filepath.Join(path, "dev"),
			filepath.Join(path, "size"),
		})
		if err != nil {
			return Info{}, err
		}

		d := Disk{}

		d.Name = filepath.Base(path)
		d.Device = o["dev"]

		d.SizeGB, err = strconv.Atoi(o["size"])
		if err != nil {
			return Info{}, err
		}
		d.SizeGB = d.SizeGB * 512 / 1024 / 1024 / 1024

		i.Disks = append(i.Disks, d)
	}

	return i, nil
}

// GetInfo return information about a systems memory.
/*
func GetInfo() (Info, error) {
	i := Info{}

	fn := "/proc/partitions"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return Info{}, fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return Info{}, err
	}

	for c, line := range strings.Split(string(o), "\n") {
		vals := strings.Fields(line)
		if c == 0 || len(vals) < 4 {
			continue
		}

		d := Disk{}

		d.Major, err = strconv.Atoi(vals[0])
		if err != nil {
			return Info{}, err
		}

		d.Minor, err = strconv.Atoi(vals[1])
		if err != nil {
			return Info{}, err
		}

		d.Blocks, err = strconv.Atoi(vals[2])
		if err != nil {
			return Info{}, err
		}

		d.Name = vals[3]

		i.Disks = append(i.Disks, d)
	}

	return i, nil
}
*/
