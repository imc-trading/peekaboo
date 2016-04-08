// +build linux

package daemon

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/imc-trading/peekaboo/docker"
	"github.com/imc-trading/peekaboo/docker/containers"
	"github.com/imc-trading/peekaboo/docker/images"
	"github.com/imc-trading/peekaboo/log"
	"github.com/imc-trading/peekaboo/network/interfaces"
	"github.com/imc-trading/peekaboo/network/routes"
	"github.com/imc-trading/peekaboo/storage/disks"
	"github.com/imc-trading/peekaboo/storage/lvm/logvols"
	"github.com/imc-trading/peekaboo/storage/lvm/physvols"
	"github.com/imc-trading/peekaboo/storage/lvm/volgrps"
	"github.com/imc-trading/peekaboo/storage/mounts"
	"github.com/imc-trading/peekaboo/system"
	"github.com/imc-trading/peekaboo/system/cpu"
	"github.com/imc-trading/peekaboo/system/ipmi"
	"github.com/imc-trading/peekaboo/system/memory"
	"github.com/imc-trading/peekaboo/system/opsys"
	"github.com/imc-trading/peekaboo/system/opsys/kernelcfg"
	"github.com/imc-trading/peekaboo/system/sysctls"
)

func New() Daemon {
	return &daemon{
		data: map[string]interface{}{},
		cache: map[string]cache{
			apiURL + "/system":               {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/os":            {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/os/kernelcfg":  {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/cpu":           {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/memory":        {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/sysctls":       {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/ipmi":          {Timeout: 5 * 60},  // 5 min.
			apiURL + "/network/interfaces":   {Timeout: 5 * 60},  // 5 min.
			apiURL + "/network/routes":       {Timeout: 5 * 60},  // 5 min.
			apiURL + "/storage/disks":        {Timeout: 5 * 60},  // 5 min.
			apiURL + "/storage/mounts":       {Timeout: 5 * 60},  // 5 min.
			apiURL + "/storage/lvm/physvols": {Timeout: 15 * 60}, // 15 min.
			apiURL + "/storage/lvm/logvols":  {Timeout: 15 * 60}, // 15 min.
			apiURL + "/storage/lvm/volgrps":  {Timeout: 15 * 60}, // 15 min.
			apiURL + "/docker":               {Timeout: 5 * 60},  // 5 min.
			apiURL + "/docker/containers":    {Timeout: 5 * 60},  // 5 min.
			apiURL + "/docker/images":        {Timeout: 5 * 60},  // 5 min.
		},
		router: mux.NewRouter(),
	}
}

func (d *daemon) Run(bind string, static string) error {
	// Add routes.
	d.addAPIRoute(apiURL+"/network/interfaces", interfaces.GetInterface)
	d.addAPIRoute(apiURL+"/network/routes", routes.GetInterface)
	d.addAPIRoute(apiURL+"/system", system.GetInterface)
	d.addAPIRoute(apiURL+"/system/os", opsys.GetInterface)
	d.addAPIRoute(apiURL+"/system/os/kernelcfg", kernelcfg.GetInterface)
	d.addAPIRoute(apiURL+"/system/cpu", cpu.GetInterface)
	d.addAPIRoute(apiURL+"/system/memory", memory.GetInterface)
	d.addAPIRoute(apiURL+"/system/sysctls", sysctls.GetInterface)
	d.addAPIRoute(apiURL+"/system/ipmi", ipmi.GetInterface)
	d.addAPIRoute(apiURL+"/storage/disks", disks.GetInterface)
	d.addAPIRoute(apiURL+"/storage/mounts", mounts.GetInterface)
	d.addAPIRoute(apiURL+"/storage/lvm/physvols", physvols.GetInterface)
	d.addAPIRoute(apiURL+"/storage/lvm/logvols", logvols.GetInterface)
	d.addAPIRoute(apiURL+"/storage/lvm/volgrps", volgrps.GetInterface)
	d.addAPIRoute(apiURL+"/docker", docker.GetInterface)
	d.addAPIRoute(apiURL+"/docker/containers", containers.GetInterface)
	d.addAPIRoute(apiURL+"/docker/images", images.GetInterface)
	d.addStaticRoute("/", static)

	logr := handlers.LoggingHandler(os.Stderr, d.router)

	log.Infof("Bind to address and port: %s", bind)
	err := http.ListenAndServe(bind, logr)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
