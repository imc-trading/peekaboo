// +build linux

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mickep76/hwinfo"

	"github.com/imc-trading/peekaboo/log"
)

type timeout struct {
	Timeout int `json:"timeout_sec"`
}

func writeJSON(w http.ResponseWriter, r *http.Request, data interface{}, cache interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if strings.ToLower(r.URL.Query().Get("envelope")) == "true" {
		e := map[string]interface{}{
			"status": http.StatusOK,
			"data":   data,
			"cache":  cache,
		}

		writeMIME(w, r, e)
	} else {
		writeMIME(w, r, data)
	}
}

func writeMIME(w http.ResponseWriter, r *http.Request, data interface{}) {
	var b []byte
	if strings.ToLower(r.URL.Query().Get("indent")) == "false" {
		b, _ = json.Marshal(data)
	} else {
		b, _ = json.MarshalIndent(data, "", "  ")
	}
	w.Write(b)
}

func htmlTemplate(title, templ string, hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if err := hwi.Update(); err != nil {
			log.Fatal(err.Error())
		}

		// Input
		input := map[string]interface{}{
			"Title":   title,
			"Version": Version,
			"HWInfo":  hwi.GetData(),
		}

		// Write template.
		b := new(bytes.Buffer)
		if err := templates.ExecuteTemplate(b, templ, input); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(b.Bytes())
	}
}

// Get
func apiGet(forceUpdate func() error, update func() error, getData func() interface{}, getCache func() interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if strings.ToLower(r.URL.Query().Get("update")) == "true" {
			if err := forceUpdate(); err != nil {
				log.Fatal(err.Error())
			}
		} else {
			if err := update(); err != nil {
				log.Fatal(err.Error())
			}
		}

		writeJSON(w, r, getData(), getCache())
	}
}

// Update
func apiUpdate(forceUpdate func() error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if err := forceUpdate(); err != nil {
			log.Fatal(err.Error())
		}
	}
}

// Set Tineout
func apiSetTimeout(setTimeout func(int)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Decode JSON into struct
		t := timeout{}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Infof("Routes timeout changed to: %d", t.Timeout)
		setTimeout(t.Timeout)
	}
}

// Add API route
func apiAddRoute(r *mux.Router, endpoint string, forceUpdate func() error, update func() error, setTimeout func(int), getData func() interface{}, getCache func() interface{}) {
	log.Infof("Add API endpoint: %s method: %s", endpoint, "GET")
	r.HandleFunc(endpoint, apiGet(forceUpdate, update, getData, getCache)).Methods("GET")

	log.Infof("Add API endpoint: %s method: %s", endpoint+"/update", "PUT")
	r.HandleFunc(endpoint+"/update", apiUpdate(forceUpdate)).Methods("PUT")

	log.Infof("Add static endpoint: %s path: %s", endpoint, "PUT")
	r.HandleFunc(endpoint, apiSetTimeout(setTimeout)).Methods("PUT")
}

// Add HTML route
func htmlAddRoute(r *mux.Router, endpoint, title, templ string, hwi hwinfo.HWInfo) {
	log.Infof("Add HTML endpoint: %s template: %s", endpoint, templ)
	r.HandleFunc(endpoint, htmlTemplate(title, templ, hwi)).Methods("GET")
}

func routes(r *mux.Router, hwi hwinfo.HWInfo) {
	// Dashboard
	htmlAddRoute(r, "/", "Dashboard", "dashboard", hwi)
	htmlAddRoute(r, "/network", "Network", "network", hwi)
	htmlAddRoute(r, "/storage", "Storage", "storage", hwi)
	htmlAddRoute(r, "/pci", "PCI", "pci", hwi)
	htmlAddRoute(r, "/sysctl", "Sysctl", "sysctl", hwi)
	htmlAddRoute(r, "/dock2box", "Dock2Box", "dock2box", hwi)

	apiURL := "/api/v1"

	// All
	//	r.HandleFunc(apiURL+"/", apiGet(hwi.Update, hwi.Update, hwi.SetTimeout, hwi.GetDataIntf, hwi.GetCacheIntf)).Methods("GET")

	// CPU
	cpu := hwi.GetCPU()
	apiAddRoute(r, apiURL+"/cpu", cpu.ForceUpdate, cpu.Update, cpu.SetTimeout, cpu.GetDataIntf, cpu.GetCacheIntf)

	// Memory
	mem := hwi.GetMemory()
	apiAddRoute(r, apiURL+"/memory", mem.ForceUpdate, mem.Update, mem.SetTimeout, mem.GetDataIntf, mem.GetCacheIntf)

	// Interfaces
	ifs := hwi.GetInterfaces()
	apiAddRoute(r, apiURL+"/interfaces", ifs.ForceUpdate, ifs.Update, ifs.SetTimeout, ifs.GetDataIntf, ifs.GetCacheIntf)

	// OpSys
	os := hwi.GetInterfaces()
	apiAddRoute(r, apiURL+"/os", os.ForceUpdate, os.Update, os.SetTimeout, os.GetDataIntf, os.GetCacheIntf)

	// Disks
	disks := hwi.GetDisks()
	apiAddRoute(r, apiURL+"/disks", disks.ForceUpdate, disks.Update, disks.SetTimeout, disks.GetDataIntf, disks.GetCacheIntf)

	// PCI
	pci := hwi.GetPCI()
	apiAddRoute(r, apiURL+"/pci", pci.ForceUpdate, pci.Update, pci.SetTimeout, pci.GetDataIntf, pci.GetCacheIntf)

	// Routes
	routes := hwi.GetRoutes()
	apiAddRoute(r, apiURL+"/routes", routes.ForceUpdate, routes.Update, routes.SetTimeout, routes.GetDataIntf, routes.GetCacheIntf)

	// PhysVols
	pvs := hwi.GetPhysVols()
	apiAddRoute(r, apiURL+"/lvm/physvols", pvs.ForceUpdate, pvs.Update, pvs.SetTimeout, pvs.GetDataIntf, pvs.GetCacheIntf)

	// LogVols
	lvs := hwi.GetLogVols()
	apiAddRoute(r, apiURL+"/lvm/logvols", lvs.ForceUpdate, lvs.Update, lvs.SetTimeout, lvs.GetDataIntf, lvs.GetCacheIntf)

	// VolGrps
	vgs := hwi.GetVolGrps()
	apiAddRoute(r, apiURL+"/lvm/volgrps", vgs.ForceUpdate, vgs.Update, vgs.SetTimeout, vgs.GetDataIntf, vgs.GetCacheIntf)

	// System
	sys := hwi.GetSystem()
	apiAddRoute(r, apiURL+"/system", sys.ForceUpdate, sys.Update, sys.SetTimeout, sys.GetDataIntf, sys.GetCacheIntf)

	// Sysctl
	sysctl := hwi.GetSysctl()
	apiAddRoute(r, apiURL+"/sysctl", sysctl.ForceUpdate, sysctl.Update, sysctl.SetTimeout, sysctl.GetDataIntf, sysctl.GetCacheIntf)

	// Dock2Box
	d2b := hwi.GetDock2Box()
	apiAddRoute(r, apiURL+"/dock2box", d2b.ForceUpdate, d2b.Update, d2b.SetTimeout, d2b.GetDataIntf, d2b.GetCacheIntf)

	// Mounts
	mnts := hwi.GetMounts()
	apiAddRoute(r, apiURL+"/mounts", mnts.ForceUpdate, mnts.Update, mnts.SetTimeout, mnts.GetDataIntf, mnts.GetCacheIntf)
}
