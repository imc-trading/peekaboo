package hwtypes

import (
	//	"encoding/json"
	"fmt"

	"github.com/mickep76/dquery"

	"github.com/imc-trading/peekaboo/docker"
	"github.com/imc-trading/peekaboo/docker/containers"
	"github.com/imc-trading/peekaboo/docker/images"
	"github.com/imc-trading/peekaboo/network"
	"github.com/imc-trading/peekaboo/network/interfaces"
	"github.com/imc-trading/peekaboo/network/routes"
	"github.com/imc-trading/peekaboo/network/arp"
	"github.com/imc-trading/peekaboo/storage/disks"
	"github.com/imc-trading/peekaboo/storage/filesystems"
	"github.com/imc-trading/peekaboo/storage/lvm/logvols"
	"github.com/imc-trading/peekaboo/storage/lvm/physvols"
	"github.com/imc-trading/peekaboo/storage/lvm/volgrps"
	"github.com/imc-trading/peekaboo/storage/mounts"
	"github.com/imc-trading/peekaboo/system"
	"github.com/imc-trading/peekaboo/system/cpu"
	"github.com/imc-trading/peekaboo/system/cpu/load"
	"github.com/imc-trading/peekaboo/system/ipmi"
	"github.com/imc-trading/peekaboo/system/ipmi/sensors"
	"github.com/imc-trading/peekaboo/system/kernel/config"
	"github.com/imc-trading/peekaboo/system/kernel/modules"
	"github.com/imc-trading/peekaboo/system/memory"
	"github.com/imc-trading/peekaboo/system/opsys"
	"github.com/imc-trading/peekaboo/system/pcicards"
	"github.com/imc-trading/peekaboo/system/rpms"
	"github.com/imc-trading/peekaboo/system/sysctls"
)

var hwTypes = []string{
	"network (short: net)",
	"network/interfaces (short: ifs)",
	"network/routes (short: routes)",
	"network/arp (short: arp)",
	"system (short: sys)",
	"system/cpu (short: cpu)",
	"system/cpu/load (short: load)",
	"system/memory (short: mem)",
	"system/os (short: os)",
	"system/kernelcfg (short: kcfg)",
	"system/sysctls (short: sysctls)",
	"system/ipmi (short: ipmi)",
	"system/ipmi/sensors (short: sensors)",
	"system/rpms (short: rpms)",
	"system/pcicards (short: pci)",
	"system/modules (short: mods)",
	"storage/disks (short: disks)",
	"storage/mounts (short: disks)",
	"storage/lvm/physvols (short: pvs)",
	"storage/lvm/logvols (short: lvs)",
	"storage/lvm/volgrps (short: vgs)",
	"storage/filesystems (short: fs)",
	"docker (short: dkr)",
	"docker/containers (short: cnts)",
	"docker/images (short: imgs)",
}

func Get(hwType string, filter string) error {
	var r interface{}
	var err error

	switch hwType {
	case "net", "network":
		r, err = network.Get()
	case "ifs", "network/interfaces":
		r, err = interfaces.GetInterface()
	case "routes", "network/routes":
		r, err = routes.Get()
	case "arp", "network/arp":
		r, err = arp.Get()
	case "sys", "system":
		r, err = system.Get()
	case "os", "system/os":
		r, err = opsys.Get()
	case "kcfg", "system/kernel/config":
		r, err = config.Get()
	case "cpu", "system/cpu":
		r, err = cpu.Get()
	case "load", "system/cpu/load":
		r, err = load.Get()
	case "sysctls", "system/sysctls":
		r, err = sysctls.Get()
	case "mem", "system/memory":
		r, err = memory.Get()
	case "ipmi", "system/ipmi":
		r, err = ipmi.Get()
	case "sensors", "system/ipmi/sensors":
		r, err = sensors.Get()
	case "rpms", "system/rpms":
		r, err = rpms.Get()
	case "pci", "system/pcicards":
		r, err = pcicards.Get()
	case "mods", "system/kernel/modules":
		r, err = modules.Get()
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
	case "fs", "storage/filesystems":
		r, err = filesystems.Get()
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

	j, err := dquery.FilterJSON(filter, r)
	if err != nil {
		return err
	}
	fmt.Println(string(j))

	return nil
}
