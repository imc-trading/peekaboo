package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/Unknwon/macaron"
	flags "github.com/jessevdk/go-flags"

	"github.com/mickep76/hwinfo"
)

func main() {
	// Set log options.
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)

	// Options.
	var opts struct {
		Verbose bool `short:"v" long:"verbose" description:"Verbose"`
		Version bool `long:"version" description:"Version"`
	}

	// Parse options.
	if _, err := flags.Parse(&opts); err != nil {
		ferr := err.(*flags.Error)
		if ferr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			log.Fatal(err.Error())
		}
	}

	// Print version.
	if opts.Version {
		fmt.Printf("go-peekaboo %s\n", Version)
		os.Exit(0)
	}

	// Set verbose.
	if opts.Verbose {
		log.SetLevel(log.InfoLevel)
	}

	info, err := hwinfo.GetInfo()
	if err != nil {
		log.Fatal(err.Error())
	}

	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/", func(ctx *macaron.Context) {
		ctx.Redirect("/system")
	})

	m.Get("/system", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "System"
		ctx.Data["CPU"] = info.CPU
		ctx.Data["Memory"] = info.Memory
		ctx.Data["OS"] = info.OS
		ctx.Data["System"] = info.System

		ctx.HTML(200, "system")
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

	m.Get("/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info)
	})

	m.Get("/cpu/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.CPU)
	})

	m.Get("/memory/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Memory)
	})

	m.Get("/os/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.OS)
	})

	m.Get("/system/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.System)
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

	m.Get("/network/routes/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Routes)
	})

	m.Get("/pci/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.PCI.PCI)
	})

	m.Get("/disks/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Disk.Disks)
	})

	m.Get("/sysctl/json", func(ctx *macaron.Context) {
		ctx.JSON(200, &info.Sysctl)
	})

	/*
	       m.Get("/pci/:bus/json", func(ctx *macaron.Context) {

	   		ctx.Params("bus")

	           d, err := pciinfo.GetInfo()
	           if err != nil {
	               log.Fatal(err.Error())
	           }

	           ctx.JSON(200, &d)
	       })
	*/

	m.Run("0.0.0.0", 8080)
}
