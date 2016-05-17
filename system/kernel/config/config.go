// +build linux

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type KernelCfgs []KernelCfg

type KernelCfg struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Get() (KernelCfgs, error) {
	kernel, err := parse.Exec("uname", []string{"-r"})
	if err != nil {
		return KernelCfgs{}, err
	}

	fn := "/boot/config-" + kernel
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return KernelCfgs{}, fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return KernelCfgs{}, err
	}

	k := KernelCfgs{}
	for _, line := range strings.Split(string(o), "\n") {
		v := strings.Split(line, "=")
		if len(v) < 2 {
			continue
		}

		k = append(k, KernelCfg{
			Key:   v[0],
			Value: strings.Trim(v[1], "\""),
		})
	}

	return k, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
