package main

import (
	"fmt"
	"os"
	"runtime"

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
		Verbose     bool   `short:"v" long:"verbose" description:"Verbose"`
		Version     bool   `long:"version" description:"Version"`
		BindAddr    string `short:"b" long:"bind-addr" description:"Bind to address" default:"0.0.0.0"`
		Port        int    `short:"p" long:"port" description:"Port" default:"5050"`
		StaticDir   string `short:"s" long:"static-dir" description:"Static content" default:"static"`
		TemplateDir string `short:"t" long:"template-dir" description:"Templates" default:"templates"`
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

	// Check root.
	if runtime.GOOS != "darinw" && os.Getuid() != 0 {
		log.Fatal("This application requires root privileges to run.")
	}

	info, err := hwinfo.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	m := macaron.Classic()
	m.Use(macaron.Static(opts.StaticDir))
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  opts.TemplateDir,
		IndentJSON: true,
	}))

	routes(m, info)
	m.Run(opts.BindAddr, opts.Port)
}
