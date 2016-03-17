// +build darwin

package main

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mickep76/hwinfo"

	"github.com/imc-trading/peekaboo/log"
)

type envelope struct {
	Data  interface{} `json:"data"`
	Cache interface{} `json:"cache"`
}

func dashboard(hwi hwinfo.HWInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Update cache
		if err := hwi.Update(); err != nil {
			log.Fatal(err.Error())
		}

		// Input
		d := hwi.GetData()
		input := map[string]interface{}{
			"Title":         "Dashboard",
			"Hostname":      d.Hostname,
			"ShortHostname": d.ShortHostname,
			"Version":       Version,
			"CPU":           d.CPU,
			"OpSys":         d.OpSys,
			"Memory":        d.Memory,
			"System":        d.System,
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
		d := hwi.GetData()
		input := map[string]interface{}{
			"Title":         "Network",
			"ShortHostname": d.ShortHostname,
			"OpSys":         d.OpSys,
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

func routes(r *mux.Router, hwi hwinfo.HWInfo) {

	log.Infof("Add endpoint: %s template: %s", "/", "dashboard")
	r.HandleFunc("/", dashboard(hwi)).Methods("GET")

	log.Infof("Add endpoint: %s template: %s", "/network", "network")
	r.HandleFunc("/network", network(hwi)).Methods("GET")

	/*

		m.Get("/network", func(ctx *macaron.Context) {
			ctx.Data["Title"] = "Network"

			// Update cache
			//		if err := hwi.Update(); err != nil {
			//			log.Fatal(err.Error())
			//		}

			d := hwi.GetData()

			ctx.Data["ShortHostname"] = d.ShortHostname
			ctx.Data["OpSys"] = d.OpSys
			ctx.HTML(200, "network")
		})

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
