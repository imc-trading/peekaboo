// +build linux

package opsys

import (
	"runtime"
	"time"

	"github.com/mickep76/hwinfo/common"
)

func (op *opSys) ForceUpdate() error {
	op.cache.LastUpdated = time.Now()
	op.cache.FromCache = false

	o, err := common.ExecCmdFields("lsb_release", []string{"-a"}, ":", []string{
		"Distributor ID",
		"Release",
	})
	if err != nil {
		return err
	}

	op.data.Kernel = runtime.GOOS
	op.data.Product = o["Distributor ID"]
	op.data.ProductVersion = o["Release"]

	op.data.KernelVersion, err = common.ExecCmd("uname", []string{"-r"})
	if err != nil {
		return err
	}

	return nil
}
