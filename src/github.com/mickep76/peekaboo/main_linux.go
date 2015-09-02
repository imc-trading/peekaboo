// +build linux

package main

import (
	"github.com/Unknwon/macaron"
	"github.com/mickep76/hwinfo"
)

func routes() {
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "Peekaboo"
		ctx.Data["Version"] = Version
		ctx.Data["Hostname"] = info.Hostname
		ctx.Data["CPU"] = info.CPU
		ctx.Data["Memory"] = info.Memory
		ctx.Data["OpSys"] = info.OpSys
		ctx.Data["System"] = info.System

		ctx.HTML(200, "peekaboo")
	})

	m.Get("/network", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "Network"
		ctx.HTML(200, "network")
	})

	m.Get("/storage", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "Storage"
		ctx.HTML(200, "storage")
	})

	m.Get("/pci", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "PCI"
		ctx.HTML(200, "pci")
	})

	m.Get("/sysctl", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "Sysctl"
		ctx.HTML(200, "sysctl")
	})

	m.Get("/disks/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Disks)
	})

	m.Get("/mounts/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Mounts)
	})

	m.Get("/network/routes/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Routes)
	})

	m.Get("/sysctl/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Sysctl)
	})

	m.Get("/lvm/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.LVM)
	})

	m.Get("/lvm/phys_vols/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.LVM.PhysVols)
	})

	m.Get("/lvm/log_vols/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.LVM.LogVols)
	})

	m.Get("/lvm/vol_grps/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.LVM.VolGrps)
	})

	m.Get("/pci/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.PCI)
	})

	m.Get("/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info)
	})

	m.Get("/cpu/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.CPU)
	})

	m.Get("/memory/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Memory)
	})

	m.Get("/network/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Network)
	})

	m.Get("/network/interfaces/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Network.Interfaces)
	})

	m.Get("/opsys/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.OpSys)
	})

	m.Get("/network/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Network)
	})
}
