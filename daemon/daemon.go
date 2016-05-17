package daemon

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/imc-trading/peekaboo/log"
)

var apiURL = "/api"

type Daemon interface {
	Run(string, string) error
}

type daemon struct {
	data   map[string]interface{}
	cache  map[string]cache
	router *mux.Router
}

type cache struct {
	LastUpdated time.Time `json:"lastUpdated"`
	Timeout     int       `json:"timeoutSec"`
	FromCache   bool      `json:"fromCache"`
}

func (d *daemon) addStaticRoute(endpoint, path string) {
	endpoint = strings.TrimRight(endpoint, "/") + "/"
	log.Infof("Add static endpoint: %s path: %s", endpoint, path)
	d.router.PathPrefix(endpoint).Handler(http.StripPrefix(endpoint, http.FileServer(http.Dir(path))))
	http.Handle(endpoint, d.router)
}

func (d *daemon) funcAPIRoute(endpoint string, getHwType func() (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Initialize or refresh cache when it has expired.
		expire := d.cache[endpoint].LastUpdated.Add(time.Duration(d.cache[endpoint].Timeout) * time.Second)
		if strings.ToLower(r.URL.Query().Get("refresh")) == "true" || d.cache[endpoint].LastUpdated.IsZero() || expire.Before(time.Now()) {
			var err error
			d.data[endpoint], err = getHwType()
			if err != nil {
				writeJSONError(w, r, d.data[endpoint], err.Error(), http.StatusInternalServerError)
				return
			}
			d.cache[endpoint] = cache{
				Timeout:     d.cache[endpoint].Timeout,
				LastUpdated: time.Now(),
				FromCache:   false,
			}

			writeJSON(w, r, d.data[endpoint], d.cache[endpoint])
			return
		}

		d.cache[endpoint] = cache{
			Timeout:     d.cache[endpoint].Timeout,
			LastUpdated: d.cache[endpoint].LastUpdated,
			FromCache:   true,
		}

		writeJSON(w, r, d.data[endpoint], d.cache[endpoint])
		return
	}
}

func (d *daemon) addAPIRoute(endpoint string, getHwType func() (interface{}, error)) {
	log.Infof("Add API endpoint: %s method: GET", endpoint)
	d.router.HandleFunc(endpoint, d.funcAPIRoute(endpoint, getHwType)).Methods("GET")

	log.Infof("Add API purge cache endpoint: %s method: PURGE", endpoint)
	d.router.HandleFunc(endpoint, d.funcAPIPurgeRoute(endpoint, getHwType)).Methods("PURGE")
}

func (d *daemon) funcAPIPurgeRoute(endpoint string, getHwType func() (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d.data[endpoint] = map[string]interface{}{}
		d.cache[endpoint] = cache{
			Timeout: d.cache[endpoint].Timeout,
		}
		writeJSON(w, r, d.data[endpoint], d.cache[endpoint])
	}
}
