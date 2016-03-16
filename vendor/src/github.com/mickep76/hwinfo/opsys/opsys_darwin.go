// +build darwin

package opsys

import (
	"runtime"
	"time"

	"github.com/mickep76/hwinfo/common"
)

func (op *opSys) ForceUpdate() error {
	op.cache.LastUpdated = time.Now()
	op.cache.FromCache = false

	o, err := common.ExecCmdFields("/usr/bin/sw_vers", []string{}, ":", []string{
		"ProductName",
		"ProductVersion",
	})
	if err != nil {
		return err
	}

	op.data.Kernel = runtime.GOOS
	op.data.Product = o["ProductName"]
	op.data.ProductVersion = o["ProductVersion"]

	op.data.KernelVersion, err = common.ExecCmd("uname", []string{"-r"})
	if err != nil {
		return err
	}

	return nil
}
