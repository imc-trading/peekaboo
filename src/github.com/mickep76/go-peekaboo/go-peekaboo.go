package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/Unknwon/macaron"
	flags "github.com/jessevdk/go-flags"
	//	"github.com/macaron-contrib/pongo2"
	"github.com/mickep76/hwinfo"
	"github.com/mickep76/hwinfo/cpuinfo"
	"github.com/mickep76/hwinfo/meminfo"
	"github.com/mickep76/hwinfo/netinfo"
	"github.com/mickep76/hwinfo/osinfo"
	"github.com/mickep76/hwinfo/pciinfo"
	"github.com/mickep76/hwinfo/sysinfo"
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

	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/system", func(ctx *macaron.Context) {
		ctx.Data["title"] = "System"
		d, err := cpuinfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}
		ctx.Data["CPU"] = d

		d2, err2 := meminfo.GetInfo()
		if err2 != nil {
			log.Fatal(err2.Error())
		}
		ctx.Data["Mem"] = d2

		d3, err3 := osinfo.GetInfo()
		if err3 != nil {
			log.Fatal(err3.Error())
		}
		ctx.Data["OS"] = d3

		d4, err4 := sysinfo.GetInfo()
		if err4 != nil {
			log.Fatal(err4.Error())
		}
		ctx.Data["Sys"] = d4

		ctx.HTML(200, "system")
	})

	m.Get("/network", func(ctx *macaron.Context) {
		ctx.Data["title"] = "Network"
		ctx.HTML(200, "network")
	})

	m.Get("/pci", func(ctx *macaron.Context) {
		ctx.Data["title"] = "PCI"
		ctx.HTML(200, "pci")
	})

	m.Get("/json", func(ctx *macaron.Context) {
		d, err := hwinfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/cpu/json", func(ctx *macaron.Context) {
		d, err := cpuinfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/mem/json", func(ctx *macaron.Context) {
		d, err := meminfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/os/json", func(ctx *macaron.Context) {
		d, err := osinfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/sys/json", func(ctx *macaron.Context) {
		d, err := sysinfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/mem/json", func(ctx *macaron.Context) {
		d, err := meminfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/network/json", func(ctx *macaron.Context) {
		d, err := netinfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/network/interfaces/json", func(ctx *macaron.Context) {
		d, err := netinfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d.Interfaces)
	})

	m.Get("/pci/json", func(ctx *macaron.Context) {
		d, err := pciinfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d.PCI)
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
