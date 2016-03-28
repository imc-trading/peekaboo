// +build linux

package opsys

import (
	"runtime"

	"github.com/imc-trading/peekaboo/parse"
)

func Get() (OpSys, error) {
	op := OpSys{}

	o, err := parse.ExecRegexpMap("lsb_release", []string{"-a"}, ":", "\\S+:\\s\\S+")
	if err != nil {
		return OpSys{}, err
	}

	op.Kernel = runtime.GOOS
	op.Product = o["Distributor ID"]
	op.ProductVersion = o["Release"]

	op.KernelVersion, err = parse.Exec("uname", []string{"-r"})
	if err != nil {
		return OpSys{}, err
	}

	return op, nil
}
