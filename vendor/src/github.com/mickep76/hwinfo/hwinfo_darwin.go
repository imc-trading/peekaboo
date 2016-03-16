package hwinfo

import (
	"os"
	"strings"

	"github.com/mickep76/hwinfo/cpu"
	"github.com/mickep76/hwinfo/interfaces"
	"github.com/mickep76/hwinfo/memory"
	"github.com/mickep76/hwinfo/opsys"
	"github.com/mickep76/hwinfo/system"
)

type HWInfo interface {
	Update() error
	GetData() data
	GetCache() cache
	GetCPU() cpu.CPU
	GetSystem() system.System
	GetMemory() memory.Memory
	GetOpSys() opsys.OpSys
	GetInterfaces() interfaces.Interfaces
}

type hwInfo struct {
	CPU        cpu.CPU
	System     system.System
	Memory     memory.Memory
	OpSys      opsys.OpSys
	Interfaces interfaces.Interfaces
	data       *data
	cache      *cache
}

type data struct {
	Hostname      string          `json:"hostname"`
	ShortHostname string          `json:"short_hostname"`
	CPU           cpu.Data        `json:"cpu"`
	System        system.Data     `json:"system"`
	Memory        memory.Data     `json:"memory"`
	OpSys         opsys.Data      `json:"opsys"`
	Interfaces    interfaces.Data `json:"interfaces"`
}

type cache struct {
	CPU        cpu.Cache        `json:"cpu"`
	System     system.Cache     `json:"system"`
	Memory     memory.Cache     `json:"memory"`
	OpSys      opsys.Cache      `json:"opsys"`
	Interfaces interfaces.Cache `json:"interfaces"`
}

func New() HWInfo {
	return &hwInfo{
		CPU:        cpu.New(),
		System:     system.New(),
		Memory:     memory.New(),
		OpSys:      opsys.New(),
		Interfaces: interfaces.New(),
		data:       &data{},
		cache:      &cache{},
	}
}

func (h *hwInfo) GetCPU() cpu.CPU {
	return h.CPU
}

func (h *hwInfo) GetSystem() system.System {
	return h.System
}

func (h *hwInfo) GetMemory() memory.Memory {
	return h.Memory
}

func (h *hwInfo) GetOpSys() opsys.OpSys {
	return h.OpSys
}

func (h *hwInfo) GetInterfaces() interfaces.Interfaces {
	return h.Interfaces
}

func (h *hwInfo) Update() error {
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	h.data.Hostname = host
	h.data.ShortHostname = strings.Split(host, ".")[0]

	if err := h.CPU.Update(); err != nil {
		return err
	}
	h.data.CPU = h.CPU.GetData()
	h.cache.CPU = h.CPU.GetCache()

	if err := h.System.Update(); err != nil {
		return err
	}
	h.data.System = h.System.GetData()
	h.cache.System = h.System.GetCache()

	if err := h.Memory.Update(); err != nil {
		return err
	}
	h.data.Memory = h.Memory.GetData()
	h.cache.Memory = h.Memory.GetCache()

	if err := h.Interfaces.Update(); err != nil {
		return err
	}
	h.data.Interfaces = h.Interfaces.GetData()
	h.cache.Interfaces = h.Interfaces.GetCache()

	if err := h.OpSys.Update(); err != nil {
		return err
	}
	h.data.OpSys = h.OpSys.GetData()
	h.cache.OpSys = h.OpSys.GetCache()

	return nil
}
