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

func htmlDashboard(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if err := hwi.Update(); err != nil {
			log.Fatal(err.Error())
		}

		// Input
		input := map[string]interface{}{
			"Title":   "Dashboard",
			"Version": Version,
			"HWInfo":  hwi.GetData(),
		}

		// Write template.
		b := new(bytes.Buffer)
		if err := templates.ExecuteTemplate(b, "dashboard", input); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(b.Bytes())
	}
}

func htmlNetwork(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if err := hwi.Update(); err != nil {
			log.Fatal(err.Error())
		}

		// Input
		input := map[string]interface{}{
			"Title":  "Network",
			"HWInfo": hwi.GetData(),
		}

		// Write template.
		b := new(bytes.Buffer)
		if err := templates.ExecuteTemplate(b, "network", input); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(b.Bytes())
	}
}

func apiGetAll(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if strings.ToLower(r.URL.Query().Get("update")) == "true" {
			//			if err := hwi.ForceUpdate(); err != nil {
			//				log.Fatal(err.Error())
			//			}
		} else {
			if err := hwi.Update(); err != nil {
				log.Fatal(err.Error())
			}
		}

		writeJSON(w, r, hwi.GetData(), hwi.GetCache())
	}
}

func apiUpdateAll(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		//        if err := hwi.ForceUpdate(); err != nil {
		//            log.Fatal(err.Error())
		//        }
	}
}

func apiGetCPU(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if strings.ToLower(r.URL.Query().Get("update")) == "true" {
			if err := hwi.GetCPU().ForceUpdate(); err != nil {
				log.Fatal(err.Error())
			}
		} else {
			if err := hwi.GetCPU().Update(); err != nil {
				log.Fatal(err.Error())
			}
		}

		writeJSON(w, r, hwi.GetCPU().GetData(), hwi.GetCPU().GetCache())
	}
}

func apiUpdateCPU(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if err := hwi.GetCPU().ForceUpdate(); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func apiSetTimeoutCPU(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Decode JSON into struct
		t := timeout{}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Infof("CPU timeout changed to: %d", t.Timeout)
		hwi.GetCPU().SetTimeout(t.Timeout)
	}
}

func apiGetMemory(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if strings.ToLower(r.URL.Query().Get("update")) == "true" {
			if err := hwi.GetMemory().ForceUpdate(); err != nil {
				log.Fatal(err.Error())
			}
		} else {
			if err := hwi.GetMemory().Update(); err != nil {
				log.Fatal(err.Error())
			}
		}

		writeJSON(w, r, hwi.GetMemory().GetData(), hwi.GetMemory().GetCache())
	}
}

func apiUpdateMemory(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if err := hwi.GetMemory().ForceUpdate(); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func apiSetTimeoutMemory(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Decode JSON into struct
		t := timeout{}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Infof("Memory timeout changed to: %d", t.Timeout)
		hwi.GetMemory().SetTimeout(t.Timeout)
	}
}

func apiGetInterfaces(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if strings.ToLower(r.URL.Query().Get("update")) == "true" {
			if err := hwi.GetInterfaces().ForceUpdate(); err != nil {
				log.Fatal(err.Error())
			}
		} else {
			if err := hwi.GetInterfaces().Update(); err != nil {
				log.Fatal(err.Error())
			}
		}

		writeJSON(w, r, hwi.GetInterfaces().GetData(), hwi.GetInterfaces().GetCache())
	}
}

func apiUpdateInterfaces(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if err := hwi.GetInterfaces().ForceUpdate(); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func apiSetTimeoutInterfaces(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Decode JSON into struct
		t := timeout{}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Infof("Interfaces timeout changed to: %d", t.Timeout)
		hwi.GetInterfaces().SetTimeout(t.Timeout)
	}
}

func apiGetOpSys(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if strings.ToLower(r.URL.Query().Get("update")) == "true" {
			if err := hwi.GetOpSys().ForceUpdate(); err != nil {
				log.Fatal(err.Error())
			}
		} else {
			if err := hwi.GetOpSys().Update(); err != nil {
				log.Fatal(err.Error())
			}
		}

		writeJSON(w, r, hwi.GetOpSys().GetData(), hwi.GetOpSys().GetCache())
	}
}

func apiUpdateOpSys(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if err := hwi.GetOpSys().ForceUpdate(); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func apiSetTimeoutOpSys(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Decode JSON into struct
		t := timeout{}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Infof("OpSys timeout changed to: %d", t.Timeout)
		hwi.GetOpSys().SetTimeout(t.Timeout)
	}
}

func routes(r *mux.Router, hwi hwinfo.HWInfo) {
	// Dashboard
	log.Infof("Add HTML endpoint: %s template: %s", "/", "dashboard")
	r.HandleFunc("/", htmlDashboard(hwi)).Methods("GET")

	// Network
	log.Infof("Add HTML endpoint: %s template: %s", "/network", "network")
	r.HandleFunc("/network", htmlNetwork(hwi)).Methods("GET")

	apiURL := "/api/v1"

	// All
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/", "GET")
	r.HandleFunc(apiURL+"/", apiGetAll(hwi)).Methods("GET")

	// All update
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/update", "PUT")
	r.HandleFunc(apiURL+"/update", apiUpdateAll(hwi)).Methods("PUT")

	// CPU
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/cpu", "GET")
	r.HandleFunc(apiURL+"/cpu", apiGetCPU(hwi)).Methods("GET")

	// CPU update
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/cpu/update", "PUT")
	r.HandleFunc(apiURL+"/cpu/update", apiUpdateCPU(hwi)).Methods("PUT")

	// CPU set timeout
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/cpu", "PUT")
	r.HandleFunc(apiURL+"/cpu", apiSetTimeoutCPU(hwi)).Methods("PUT")

	// Memory
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/memory", "GET")
	r.HandleFunc(apiURL+"/memory", apiGetMemory(hwi)).Methods("GET")

	// Memory update
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/memory/update", "PUT")
	r.HandleFunc(apiURL+"/memory/update", apiUpdateMemory(hwi)).Methods("PUT")

	// Memory set timeout
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/memory", "PUT")
	r.HandleFunc(apiURL+"/memory", apiSetTimeoutMemory(hwi)).Methods("PUT")

	// Interfaces
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/interfaces", "GET")
	r.HandleFunc(apiURL+"/interfaces", apiGetInterfaces(hwi)).Methods("GET")

	// Interfaces update
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/interfaces/update", "PUT")
	r.HandleFunc(apiURL+"/interfaces/update", apiUpdateInterfaces(hwi)).Methods("PUT")

	// Interfaces set timeout
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/interfaces", "PUT")
	r.HandleFunc(apiURL+"/interfaces", apiSetTimeoutInterfaces(hwi)).Methods("PUT")

	// OpSys
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/opsys", "GET")
	r.HandleFunc(apiURL+"/opsys", apiGetOpSys(hwi)).Methods("GET")

	// OpSys update
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/opsys/update", "PUT")
	r.HandleFunc(apiURL+"/opsys/update", apiUpdateOpSys(hwi)).Methods("PUT")

	// OpSys set timeout
	log.Infof("Add API endpoint: %s%s method: %s", apiURL, "/opsys", "PUT")
	r.HandleFunc(apiURL+"/opsys", apiSetTimeoutOpSys(hwi)).Methods("PUT")
}
