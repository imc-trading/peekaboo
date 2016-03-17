// +build darwin

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

func dashboard(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
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

func network(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
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

func allJSON(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if err := hwi.Update(); err != nil {
			log.Fatal(err.Error())
		}

		writeJSON(w, r, hwi.GetData(), hwi.GetCache())
	}
}

func routes(r *mux.Router, hwi hwinfo.HWInfo) {

	log.Infof("Add endpoint: %s template: %s", "/", "dashboard")
	r.HandleFunc("/", dashboard(hwi)).Methods("GET")

	log.Infof("Add endpoint: %s template: %s", "/network", "network")
	r.HandleFunc("/network", network(hwi)).Methods("GET")

	log.Infof("Add API endpoint: %s", "/json")
	r.HandleFunc("/json", allJSON(hwi)).Methods("GET")

	/*
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
	*/

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
	/*
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
	*/
}
