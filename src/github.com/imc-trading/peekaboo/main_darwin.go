// +build darwin

package main

import (
	"log"

	"github.com/Unknwon/macaron"
	"github.com/mickep76/hwinfo"
)

type envelope struct {
	Data  interface{} `json:"data"`
	Cache interface{} `json:"cache"`
}

func routes(m *macaron.Macaron, hwi hwinfo.HWInfo) {
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "Peekaboo"

		// Update cache
		if err := hwi.Update(); err != nil {
			log.Fatal(err.Error())
		}

		d := hwi.GetData()

		ctx.Data["Hostname"] = d.Hostname
		ctx.Data["ShortHostname"] = d.ShortHostname
		ctx.Data["Version"] = Version

		ctx.Data["CPU"] = d.CPU
		ctx.Data["OpSys"] = d.OpSys
		ctx.Data["Memory"] = d.Memory
		ctx.Data["System"] = d.System

		ctx.HTML(200, "peekaboo")
	})

	m.Get("/network", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "Network"

		// Update cache
		//		if err := hwi.Update(); err != nil {
		//			log.Fatal(err.Error())
		//		}

		d := hwi.GetData()

		ctx.Data["ShortHostname"] = d.ShortHostname
		ctx.Data["OpSys"] = d.OpSys
		ctx.HTML(200, "network")
	})

	m.Get("/json", func(ctx *macaron.Context) {
		// Update cache
		if err := hwi.Update(); err != nil {
			log.Fatal(err.Error())
		}

		d := hwi.GetData()
		c := hwi.GetCache()

		e := envelope{
			Data:  &d,
			Cache: &c,
		}

		ctx.JSON(200, &e)
	})

	m.Get("/cpu/json", func(ctx *macaron.Context) {
		// Update cache
		if err := hwi.GetCPU().Update(); err != nil {
			log.Fatal(err.Error())
		}

		d := hwi.GetData()

		ctx.JSON(200, &d.CPU)
	})

	m.Get("/memory/json", func(ctx *macaron.Context) {
		// Update cache
		if err := hwi.GetMemory().Update(); err != nil {
			log.Fatal(err.Error())
		}

		d := hwi.GetData()

		ctx.JSON(200, &d.Memory)
	})

	/*
		m.Get("/network/json", func(ctx *macaron.Context) {
			// Update cache
			if err := hwi.Update(); err != nil {
				log.Fatal(err.Error())
			}

			d := hwi.GetData()

			ctx.JSON(200, &d.Network)
		})
	*/

	m.Get("/network/interfaces/json", func(ctx *macaron.Context) {
		// Update cache
		if err := hwi.GetInterfaces().Update(); err != nil {
			log.Fatal(err.Error())
		}

		d := hwi.GetData()

		ctx.JSON(200, &d.Interfaces)
	})

	m.Get("/opsys/json", func(ctx *macaron.Context) {
		// Update cache
		if err := hwi.GetOpSys().Update(); err != nil {
			log.Fatal(err.Error())
		}

		d := hwi.GetData()

		ctx.JSON(200, &d.OpSys)
	})
}
