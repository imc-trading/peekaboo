// +build darwin

package opsys

import (
	"runtime"

	"github.com/imc-trading/peekaboo/parse"
)

func Get() (OpSys, error) {
	o := OpSys{}

	m, err := parse.ExecRegexpMap("/usr/bin/sw_vers", []string{}, ":", "\\S+:\\s\\S+")
	if err != nil {
		return OpSys{}, err
	}

	o.Kernel = runtime.GOOS
	o.Product = m["ProductName"]
	o.ProductVersion = m["ProductVersion"]

	o.KernelVersion, err = parse.Exec("uname", []string{"-r"})
	if err != nil {
		return OpSys{}, err
	}

	return o, nil
}
