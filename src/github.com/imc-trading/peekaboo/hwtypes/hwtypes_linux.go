package hwtypes

import (
	"encoding/json"
	"fmt"

	"github.com/imc-trading/peekaboo/docker"
	"github.com/imc-trading/peekaboo/docker/containers"
	"github.com/imc-trading/peekaboo/docker/images"
	"github.com/imc-trading/peekaboo/network/interfaces"
	"github.com/imc-trading/peekaboo/network/routes"
	"github.com/imc-trading/peekaboo/storage/disks"
	"github.com/imc-trading/peekaboo/storage/lvm/logvols"
	"github.com/imc-trading/peekaboo/storage/lvm/physvols"
	"github.com/imc-trading/peekaboo/storage/lvm/volgrps"
	"github.com/imc-trading/peekaboo/storage/mounts"
	"github.com/imc-trading/peekaboo/system"
	"github.com/imc-trading/peekaboo/system/cpu"
	"github.com/imc-trading/peekaboo/system/ipmi"
	"github.com/imc-trading/peekaboo/system/memory"
	"github.com/imc-trading/peekaboo/system/opsys"
	"github.com/imc-trading/peekaboo/system/sysctls"
)

var hwTypes = []string{
	"network/interfaces (short: ifs)",
	"network/routes (short: routes)",
	"system (short: sys)",
	"system/cpu (short: cpu)",
	"system/memory (short: mem)",
	"system/os (short: os)",
	"system/sysctls (short: sysctls)",
	"system/ipmi (short: ipmi)",
	"storage/disks (short: disks)",
	"storage/mounts (short: disks)",
	"storage/lvm/physvols (short: pvs)",
	"storage/lvm/logvols (short: lvs)",
	"storage/lvm/volgrps (short: vgs)",
	"docker (short: dkr)",
	"docker/containers (short: cnts)",
	"docker/images (short: imgs)",
}

func Get(hwType string) error {
	var r interface{}
	var err error

	switch hwType {
	case "ifs", "network/interfaces":
		r, err = interfaces.Get()
	case "routes", "network/routes":
		r, err = routes.Get()
	case "sys", "system":
		r, err = system.Get()
	case "os", "system/os":
		r, err = opsys.Get()
	case "cpu", "system/cpu":
		r, err = cpu.Get()
	case "sysctls", "system/sysctls":
		r, err = sysctls.Get()
	case "mem", "system/memory":
		r, err = memory.Get()
	case "ipmi", "system/ipmi":
		r, err = ipmi.Get()
	case "disks", "storage/disks":
		r, err = disks.Get()
	case "mounts", "storage/mounts":
		r, err = mounts.Get()
	case "pvs", "storage/lvm/physvols":
		r, err = physvols.Get()
	case "lvs", "storage/lvm/logvols":
		r, err = logvols.Get()
	case "vgs", "storage/lvm/volgrps":
		r, err = volgrps.Get()
	case "dkr", "docker":
		r, err = docker.Get()
	case "cnts", "docker/containers":
		r, err = containers.Get()
	case "imgs", "docker/images":
		r, err = images.Get()
	default:
		return fmt.Errorf("Unknown hardware type: %s", hwType)
	}

	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(b))

	return nil
}
