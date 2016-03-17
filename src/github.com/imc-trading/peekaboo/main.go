package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	flags "github.com/jessevdk/go-flags"
	"github.com/mickep76/hwinfo"

	"github.com/imc-trading/peekaboo/log"
)

var templates *template.Template

var funcs = template.FuncMap{}

func addStaticRoute(r *mux.Router, endpoint, path string) {
	log.Infof("Add endpoint: %s path: %s", endpoint, path)

	r.PathPrefix(endpoint + "/").Handler(http.StripPrefix(endpoint+"/", http.FileServer(http.Dir(path))))
	http.Handle(endpoint+"/", r)
}

func addTemplateRoute(r *mux.Router, endpoint, templ string) {
	log.Infof("Add endpoint: %s template: %s", endpoint, templ)

	r.HandleFunc(endpoint, execTemplate(templ)).Methods("GET")
}

func execTemplate(templ string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		input := map[string]interface{}{
			"vars":   mux.Vars(r),
			"params": r.URL.Query(),
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

func main() {
	// Options.
	var opts struct {
		Verbose      bool    `short:"v" long:"verbose" description:"Verbose"`
		Version      bool    `long:"version" description:"Version"`
		Bind         string  `short:"b" long:"bind" description:"Bind to address and port" default:"0.0.0.0:5050"`
		StaticDir    string  `short:"s" long:"static-dir" description:"Static content" default:"static"`
		TemplateDir  string  `short:"t" long:"template-dir" description:"Templates" default:"templates"`
		KafkaEnabled bool    `short:"K" long:"kafka" description:"Enable Kafka message bus"`
		KafkaTopic   string  `long:"kafka-topic" description:"Kafka topic" default:"peekaboo"`
		KafkaPeers   *string `long:"kafka-peers" description:"Comma-delimited list of Kafka brokers"`
		KafkaCert    *string `long:"kafka-cert" description:"Certificate file for client authentication"`
		KafkaKey     *string `long:"kafka-key" description:"Key file for client client authentication"`
		KafkaCA      *string `long:"kafka-ca" description:"CA file for TLS client authentication"`
		KafkaVerify  bool    `long:"kafka-verify" description:"Verify SSL certificate"`
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
		fmt.Printf("peekaboo %s\n", Version)
		os.Exit(0)
	}

	// Set verbose.
	if opts.Verbose {
		log.SetDebug()
	}

	// Check root.
	if runtime.GOOS != "darwin" && os.Getuid() != 0 {
		log.Fatal("This application requires root privileges to run.")
	}

	// Compile templates
	templates = template.Must(template.New("main").Funcs(funcs).ParseGlob(opts.TemplateDir + "/*.tmpl"))

	// Get hardware info.
	hwi := hwinfo.New()
	if err := hwi.Update(); err != nil {
		log.Fatal(err.Error())
	}

	log.Infof("Using static dir: %s", opts.StaticDir)
	log.Infof("Using template dir: %s", opts.TemplateDir)

	r := mux.NewRouter()
	addStaticRoute(r, "/img", opts.StaticDir+"/img")
	addStaticRoute(r, "/js", opts.StaticDir+"/js")
	addStaticRoute(r, "/bootstrap", opts.StaticDir+"/bootstrap")
	addStaticRoute(r, "/css", opts.StaticDir+"/css")

	routes(r, hwi)

	logr := handlers.LoggingHandler(os.Stderr, r)

	log.Infof("Bind to address and port: %s", opts.Bind)
	err := http.ListenAndServe(opts.Bind, logr)
	if err != nil {
		log.Fatal(err.Error())
	}
}
