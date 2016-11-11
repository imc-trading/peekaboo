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
	"github.com/imc-trading/peekaboo/network"
	"github.com/imc-trading/peekaboo/network/interfaces"
	"github.com/imc-trading/peekaboo/network/routes"
	"github.com/imc-trading/peekaboo/storage/disks"
	"github.com/imc-trading/peekaboo/storage/lvm/logvols"
	"github.com/imc-trading/peekaboo/storage/lvm/physvols"
	"github.com/imc-trading/peekaboo/storage/lvm/volgrps"
	"github.com/imc-trading/peekaboo/storage/filesystems"
	"github.com/imc-trading/peekaboo/storage/mounts"
	"github.com/imc-trading/peekaboo/system"
	"github.com/imc-trading/peekaboo/system/cpu"
	"github.com/imc-trading/peekaboo/system/ipmi"
	"github.com/imc-trading/peekaboo/system/ipmi/sensors"
	"github.com/imc-trading/peekaboo/system/kernel/config"
	"github.com/imc-trading/peekaboo/system/kernel/modules"
	"github.com/imc-trading/peekaboo/system/memory"
	"github.com/imc-trading/peekaboo/system/opsys"
	"github.com/imc-trading/peekaboo/system/pcicards"
	"github.com/imc-trading/peekaboo/system/rpms"
	"github.com/imc-trading/peekaboo/system/sysctls"
)

func New() Daemon {
	return &daemon{
		data: map[string]interface{}{},
		cache: map[string]cache{
			apiURL + "/system":                {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/os":             {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/kernel/config":  {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/cpu":            {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/memory":         {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/sysctls":        {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/ipmi":           {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/ipmi/sensors":   {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/rpms":           {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/pcicards":       {Timeout: 5 * 60},  // 5 min.
			apiURL + "/system/kernel/modules": {Timeout: 5 * 60},  // 5 min.
			apiURL + "/network":               {Timeout: 5 * 60},  // 5 min.
			apiURL + "/network/interfaces":    {Timeout: 5 * 60},  // 5 min.
			apiURL + "/network/routes":        {Timeout: 5 * 60},  // 5 min.
			apiURL + "/network/arp":           {Timeout: 5 * 60},  // 5 min.
			apiURL + "/storage/disks":         {Timeout: 5 * 60},  // 5 min.
			apiURL + "/storage/mounts":        {Timeout: 5 * 60},  // 5 min.
			apiURL + "/storage/lvm/physvols":  {Timeout: 15 * 60}, // 15 min.
			apiURL + "/storage/lvm/logvols":   {Timeout: 15 * 60}, // 15 min.
			apiURL + "/storage/lvm/volgrps":   {Timeout: 15 * 60}, // 15 min.
			apiURL + "/storage/filesystems":   {Timeout: 5 * 60}, // 5 min.
			apiURL + "/docker":                {Timeout: 5 * 60},  // 5 min.
			apiURL + "/docker/containers":     {Timeout: 5 * 60},  // 5 min.
			apiURL + "/docker/images":         {Timeout: 5 * 60},  // 5 min.
		},
		router: mux.NewRouter(),
	}
}

func (d *daemon) Run(bind string, static string) error {
	// Add routes.
	d.addAPIRoute(apiURL+"/network", network.GetInterface)
	d.addAPIRoute(apiURL+"/network/interfaces", interfaces.GetInterface)
	d.addAPIRoute(apiURL+"/network/routes", routes.GetInterface)
	d.addAPIRoute(apiURL+"/network/arp", routes.GetInterface)
	d.addAPIRoute(apiURL+"/system", system.GetInterface)
	d.addAPIRoute(apiURL+"/system/os", opsys.GetInterface)
	d.addAPIRoute(apiURL+"/system/kernel/config", config.GetInterface)
	d.addAPIRoute(apiURL+"/system/cpu", cpu.GetInterface)
	d.addAPIRoute(apiURL+"/system/memory", memory.GetInterface)
	d.addAPIRoute(apiURL+"/system/sysctls", sysctls.GetInterface)
	d.addAPIRoute(apiURL+"/system/ipmi", ipmi.GetInterface)
	d.addAPIRoute(apiURL+"/system/ipmi/sensors", sensors.GetInterface)
	d.addAPIRoute(apiURL+"/system/rpms", rpms.GetInterface)
	d.addAPIRoute(apiURL+"/system/pcicards", pcicards.GetInterface)
	d.addAPIRoute(apiURL+"/system/kernel/modules", modules.GetInterface)
	d.addAPIRoute(apiURL+"/storage/disks", disks.GetInterface)
	d.addAPIRoute(apiURL+"/storage/mounts", mounts.GetInterface)
	d.addAPIRoute(apiURL+"/storage/lvm/physvols", physvols.GetInterface)
	d.addAPIRoute(apiURL+"/storage/lvm/logvols", logvols.GetInterface)
	d.addAPIRoute(apiURL+"/storage/lvm/volgrps", volgrps.GetInterface)
	d.addAPIRoute(apiURL+"/storage/filesystems", filesystems.GetInterface)
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
