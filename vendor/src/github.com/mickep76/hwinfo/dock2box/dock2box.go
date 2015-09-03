package dock2box

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Layer struct {
	Layer   string `json:"layer"`
	Image   string `json:"image"`
	Commit  string `json:"commit"`
	Created string `json:"created"`
	Log     string `json:"logs"`
}

type Dock2Box struct {
	FirstBoot string  `json:"firstboot"`
	CFlags    string  `json:"cflags_native"`
	Layers    []Layer `json:"layers"`
}

// Get information about Dock2Box image layers.
func Get() (Dock2Box, error) {
	file := "/etc/dock2box/firstboot.json"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return Dock2Box{}, fmt.Errorf("file doesn't exist: %s", file)
	}

	o, err := ioutil.ReadFile(file)
	if err != nil {
		return Dock2Box{}, err
	}

	d2b := Dock2Box{}
	if err := json.Unmarshal(o, &d2b); err != nil {
		return Dock2Box{}, err
	}

	files, err := filepath.Glob("/etc/dock2box/layers/*.json")
	if err != nil {
		return Dock2Box{}, err
	}

	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return Dock2Box{}, fmt.Errorf("file doesn't exist: %s", file)
		}

		o, err := ioutil.ReadFile(file)
		if err != nil {
			return Dock2Box{}, err
		}

		layer := Layer{}
		if err := json.Unmarshal(o, &layer); err != nil {
			return Dock2Box{}, err
		}

		fn := path.Base(file)
		layer.Layer = strings.TrimSuffix(fn, filepath.Ext(fn))

		d2b.Layers = append(d2b.Layers, layer)
	}

	return d2b, nil
}
