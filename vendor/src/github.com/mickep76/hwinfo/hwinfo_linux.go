package hwinfo

import (
	"os"
	"strings"

	"github.com/mickep76/hwinfo/cpu"
	"github.com/mickep76/hwinfo/disks"
	"github.com/mickep76/hwinfo/dock2box"
	"github.com/mickep76/hwinfo/dock2box/layers"
	"github.com/mickep76/hwinfo/docker"
	"github.com/mickep76/hwinfo/docker/containers"
	"github.com/mickep76/hwinfo/docker/images"
	"github.com/mickep76/hwinfo/interfaces"
	"github.com/mickep76/hwinfo/lvm/logvols"
	"github.com/mickep76/hwinfo/lvm/physvols"
	"github.com/mickep76/hwinfo/lvm/volgrps"
	"github.com/mickep76/hwinfo/memory"
	"github.com/mickep76/hwinfo/mounts"
	"github.com/mickep76/hwinfo/opsys"
	"github.com/mickep76/hwinfo/pci"
	"github.com/mickep76/hwinfo/routes"
	"github.com/mickep76/hwinfo/sysctl"
	"github.com/mickep76/hwinfo/system"
)

type HWInfo interface {
	Update() error
	GetData() Data
	GetCache() Cache
	GetCPU() cpu.CPU
	GetDisks() disks.Disks
	GetDock2Box() dock2box.Dock2Box
	GetDocker() docker.Docker
	GetContainers() containers.Containers
	GetImages() images.Images
	GetLayers() layers.Layers
	GetInterfaces() interfaces.Interfaces
	GetPhysVols() physvols.PhysVols
	GetLogVols() logvols.LogVols
	GetVolGrps() volgrps.VolGrps
	GetMemory() memory.Memory
	GetMounts() mounts.Mounts
	GetOpSys() opsys.OpSys
	GetPCI() pci.PCI
	GetRoutes() routes.Routes
	GetSysctl() sysctl.Sysctl
	GetSystem() system.System
}

type hwInfo struct {
	CPU        cpu.CPU
	Disks      disks.Disks
	Dock2Box   dock2box.Dock2Box
	Docker     docker.Docker
	Containers containers.Containers
	Images     images.Images
	Layers     layers.Layers
	Interfaces interfaces.Interfaces
	PhysVols   physvols.PhysVols
	LogVols    logvols.LogVols
	VolGrps    volgrps.VolGrps
	Memory     memory.Memory
	Mounts     mounts.Mounts
	OpSys      opsys.OpSys
	PCI        pci.PCI
	Routes     routes.Routes
	Sysctl     sysctl.Sysctl
	System     system.System
	data       *Data
	cache      *Cache
}

type Data struct {
	Hostname      string          `json:"hostname"`
	ShortHostname string          `json:"short_hostname"`
	CPU           cpu.Data        `json:"cpu"`
	Disks         disks.Data      `json:"disks"`
	Dock2Box      dock2box.Data   `json:"dock2box"`
	Docker        docker.Data     `json:"docker"`
	Containers    containers.Data `json:"containers"`
	Images        images.Data     `json:"images"`
	Layers        layers.Data     `json:"layers"`
	Interfaces    interfaces.Data `json:"interfaces"`
	PhysVols      physvols.Data   `json:"phys_vols"`
	LogVols       logvols.Data    `json:"log_vols"`
	VolGrps       volgrps.Data    `json:"vol_grps"`
	Memory        memory.Data     `json:"memory"`
	Mounts        mounts.Data     `json:"mounts"`
	OpSys         opsys.Data      `json:"opsys"`
	PCI           pci.Data        `json:"pci"`
	Routes        routes.Data     `json:"routes"`
	Sysctl        sysctl.Data     `json:"sysctl"`
	System        system.Data     `json:"system"`
}

type Cache struct {
	CPU        cpu.Cache        `json:"cpu"`
	Disks      disks.Cache      `json:"disks"`
	Dock2Box   dock2box.Cache   `json:"dock2box"`
	Docker     docker.Cache     `json:"docker"`
	Containers containers.Cache `json:"containers"`
	Images     images.Images    `json:"images"`
	Layers     layers.Cache     `json:"layers"`
	Interfaces interfaces.Cache `json:"interfaces"`
	PhysVols   physvols.Cache   `json:"phys_vols"`
	LogVols    logvols.Cache    `json:"log_vols"`
	VolGrps    volgrps.Cache    `json:"vol_grps"`
	Memory     memory.Cache     `json:"memory"`
	Mounts     mounts.Cache     `json:"mounts"`
	OpSys      opsys.Cache      `json:"opsys"`
	PCI        pci.Cache        `json:"pci"`
	Routes     routes.Cache     `json:"routes"`
	Sysctl     sysctl.Cache     `json:"sysctl"`
	System     system.Cache     `json:"system"`
}

func New() HWInfo {
	return &hwInfo{
		CPU:        cpu.New(),
		Disks:      disks.New(),
		Dock2Box:   dock2box.New(),
		Docker:     docker.New(),
		Containers: containers.New(),
		Images:     images.New(),
		Layers:     layers.New(),
		Interfaces: interfaces.New(),
		PhysVols:   physvols.New(),
		LogVols:    logvols.New(),
		VolGrps:    volgrps.New(),
		Memory:     memory.New(),
		Mounts:     mounts.New(),
		OpSys:      opsys.New(),
		PCI:        pci.New(),
		Routes:     routes.New(),
		Sysctl:     sysctl.New(),
		System:     system.New(),
		data:       &Data{},
		cache:      &Cache{},
	}
}

func (h *hwInfo) GetCPU() cpu.CPU {
	return h.CPU
}

func (h *hwInfo) GetDisks() disks.Disks {
	return h.Disks
}

func (h *hwInfo) GetDock2Box() dock2box.Dock2Box {
	return h.Dock2Box
}

func (h *hwInfo) GetDocker() docker.Docker {
	return h.Docker
}

func (h *hwInfo) GetContainers() containers.Containers {
	return h.Containers
}

func (h *hwInfo) GetImages() images.Images {
	return h.Images
}

func (h *hwInfo) GetLayers() layers.Layers {
	return h.Layers
}

func (h *hwInfo) GetInterfaces() interfaces.Interfaces {
	return h.Interfaces
}

func (h *hwInfo) GetPhysVols() physvols.PhysVols {
	return h.PhysVols
}

func (h *hwInfo) GetLogVols() logvols.LogVols {
	return h.LogVols
}

func (h *hwInfo) GetVolGrps() volgrps.VolGrps {
	return h.VolGrps
}

func (h *hwInfo) GetMemory() memory.Memory {
	return h.Memory
}

func (h *hwInfo) GetMounts() mounts.Mounts {
	return h.Mounts
}

func (h *hwInfo) GetOpSys() opsys.OpSys {
	return h.OpSys
}

func (h *hwInfo) GetPCI() pci.PCI {
	return h.PCI
}

func (h *hwInfo) GetRoutes() routes.Routes {
	return h.Routes
}

func (h *hwInfo) GetSysctl() sysctl.Sysctl {
	return h.Sysctl
}

func (h *hwInfo) GetSystem() system.System {
	return h.System
}

func (h *hwInfo) Update() error {
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	h.data.Hostname = host
	h.data.ShortHostname = strings.Split(host, ".")[0]

	// CPU
	if err := h.CPU.Update(); err != nil {
		return err
	}
	h.data.CPU = h.CPU.GetData()
	h.cache.CPU = h.CPU.GetCache()

	// System
	if err := h.System.Update(); err != nil {
		return err
	}
	h.data.System = h.System.GetData()
	h.cache.System = h.System.GetCache()

	// Memory
	if err := h.Memory.Update(); err != nil {
		return err
	}
	h.data.Memory = h.Memory.GetData()
	h.cache.Memory = h.Memory.GetCache()

	// Interfaces
	if err := h.Interfaces.Update(); err != nil {
		return err
	}
	h.data.Interfaces = h.Interfaces.GetData()
	h.cache.Interfaces = h.Interfaces.GetCache()

	// OpSys
	if err := h.OpSys.Update(); err != nil {
		return err
	}
	h.data.OpSys = h.OpSys.GetData()
	h.cache.OpSys = h.OpSys.GetCache()

	// Disks
	if err := h.Disks.Update(); err != nil {
		return err
	}
	h.data.Disks = h.Disks.GetData()
	h.cache.Disks = h.Disks.GetCache()

	// Dock2Box
	if err := h.Dock2Box.Update(); err != nil {
		return err
	}
	h.data.Dock2Box = h.Dock2Box.GetData()
	h.cache.Dock2Box = h.Dock2Box.GetCache()

	// Docker
	if err := h.Docker.Update(); err != nil {
		return err
	}
	h.data.Docker = h.Docker.GetData()
	h.cache.Docker = h.Docker.GetCache()

	// Containers
	if err := h.Containers.Update(); err != nil {
		return err
	}
	h.data.Containers = h.Containers.GetData()
	h.cache.Containers = h.Containers.GetCache()

	// Mounts
	if err := h.Mounts.Update(); err != nil {
		return err
	}
	h.data.Mounts = h.Mounts.GetData()
	h.cache.Mounts = h.Mounts.GetCache()

	// Sysctl
	if err := h.Sysctl.Update(); err != nil {
		return err
	}
	h.data.Sysctl = h.Sysctl.GetData()
	h.cache.Sysctl = h.Sysctl.GetCache()

	// PCI
	if err := h.PCI.Update(); err != nil {
		return err
	}
	h.data.PCI = h.PCI.GetData()
	h.cache.PCI = h.PCI.GetCache()

	// Routes
	if err := h.Routes.Update(); err != nil {
		return err
	}
	h.data.Routes = h.Routes.GetData()
	h.cache.Routes = h.Routes.GetCache()

	// PhysVols
	if err := h.PhysVols.Update(); err != nil {
		return err
	}
	h.data.PhysVols = h.PhysVols.GetData()
	h.cache.PhysVols = h.PhysVols.GetCache()

	// LogVols
	if err := h.LogVols.Update(); err != nil {
		return err
	}
	h.data.LogVols = h.LogVols.GetData()
	h.cache.LogVols = h.LogVols.GetCache()

	// VolGroups
	if err := h.VolGrps.Update(); err != nil {
		return err
	}
	h.data.VolGrps = h.VolGrps.GetData()
	h.cache.VolGrps = h.VolGrps.GetCache()

	return nil
}
