package hwtypes

import (
	"encoding/json"
	"fmt"

	"github.com/imc-trading/peekaboo/docker"
	"github.com/imc-trading/peekaboo/docker/containers"
	"github.com/imc-trading/peekaboo/docker/images"
	"github.com/imc-trading/peekaboo/network/interfaces"
	"github.com/imc-trading/peekaboo/network/routes"
	"github.com/imc-trading/peekaboo/system"
	"github.com/imc-trading/peekaboo/system/cpu"
	"github.com/imc-trading/peekaboo/system/memory"
	"github.com/imc-trading/peekaboo/system/opsys"
	"github.com/mickep76/dquery"
)

var hwTypes = []string{
	"network/interfaces (short: ifs)",
	"network/routes (short: routes)",
	"system (short: sys)",
	"system/cpu (short: cpu)",
	"system/memory (short: mem)",
	"system/os (short: os)",
	"docker (short: dkr)",
	"docker/containers (short: cnts)",
	"docker/images (short: imgs)",
}

func Get(hwType string, filter string) error {
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
	case "mem", "system/memory":
		r, err = memory.Get()
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
	j, err := dquery.FilterJSON(filter, b)
	if err != nil {
		return err
	}
	fmt.Println(string(j))

	return nil
}
