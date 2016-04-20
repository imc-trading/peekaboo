package main

import (
	"os"
	"runtime"

	"github.com/docopt/docopt-go"

	"github.com/imc-trading/peekaboo/daemon"
	"github.com/imc-trading/peekaboo/hwtypes"
	"github.com/imc-trading/peekaboo/log"
	"github.com/imc-trading/peekaboo/version"
)

func main() {
	usage := `Peekaboo

Usage:
  peekaboo daemon [--debug] [--bind=<addr>] [--static=<dir>]
  peekaboo list
  peekaboo get <hardware-type> [--filter=<path>]
  peekaboo -h | --help
  peekaboo --version

Commands:
  daemon               Start as a daemon serving HTTP requests.
  list                 List hardware names available.
  get                  Return information about hardware.

Arguments:
  hardware-type        Name of hardware to return information about.

Options:
  -h --help            Show this screen.
  --version            Show version.
  -d --debug           Debug.
  -b --bind=<addr>     Bind to address and port. [default: 0.0.0.0:5050]
  -s --static=<dir>    Directory for static content. [default: static]
  -f --filter=<path>   Filter result. [default: .]
`

	args, err := docopt.Parse(usage, nil, true, "Peekaboo "+version.Version, false)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Set verbose.
	if args["--debug"].(bool) {
		log.SetDebug()
	}

	// Check root.
	if runtime.GOOS != "darwin" && os.Getuid() != 0 {
		log.Fatal("This application requires root privileges to run.")
	}

	// List hardware types.
	if args["list"].(bool) {
		hwtypes.List()
	}

	// Get hardware type.
	if args["get"].(bool) {
		if err := hwtypes.Get(args["<hardware-type>"].(string), args["--filter"].(string)); err != nil {
			log.Fatal(err.Error())
		}
	}

	// Run as a daemon.
	if args["daemon"].(bool) {
		d := daemon.New()
		if err := d.Run(args["--bind"].(string), args["--static"].(string)); err != nil {
			log.Fatal(err.Error())
		}
	}
}
