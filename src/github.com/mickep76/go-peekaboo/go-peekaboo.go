package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/Unknwon/macaron"
	flags "github.com/jessevdk/go-flags"
	"github.com/mickep76/hwinfo"
	"github.com/mickep76/hwinfo/cpu"
	"github.com/mickep76/hwinfo/mem"
	"github.com/mickep76/hwinfo/netinfo"
	hwos "github.com/mickep76/hwinfo/os"
	"github.com/mickep76/hwinfo/sys"
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

	m.Get("/json", func(ctx *macaron.Context) {
		d, err := hwinfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/cpu/json", func(ctx *macaron.Context) {
		d, err := cpu.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/mem/json", func(ctx *macaron.Context) {
		d, err := mem.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/os/json", func(ctx *macaron.Context) {
		d, err := hwos.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/sys/json", func(ctx *macaron.Context) {
		d, err := sys.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/mem/json", func(ctx *macaron.Context) {
		d, err := mem.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Get("/net/json", func(ctx *macaron.Context) {
		d, err := netinfo.GetInfo()
		if err != nil {
			log.Fatal(err.Error())
		}

		ctx.JSON(200, &d)
	})

	m.Run("0.0.0.0", 8080)
}
