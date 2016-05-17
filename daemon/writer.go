package daemon

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mickep76/dquery"
)

func writeJSON(w http.ResponseWriter, r *http.Request, data interface{}, cache interface{}) {
	if strings.ToLower(r.URL.Query().Get("envelope")) == "true" {
		e := map[string]interface{}{
			"status": http.StatusOK,
			"data":   data,
			"error":  []string{},
			"cache":  cache,
		}

		filter := r.URL.Query().Get("filter")
		if filter != "" {
			d, err := dquery.Filter(filter, e)
			if err != nil {
				writeJSONError(w, r, nil, err.Error(), http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			writeMIME(w, r, d)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		writeMIME(w, r, e)
	} else {

		filter := r.URL.Query().Get("filter")
		if filter != "" {
			d, err := dquery.Filter(filter, data)
			if err != nil {
				writeJSONError(w, r, nil, err.Error(), http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			writeMIME(w, r, d)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		writeMIME(w, r, data)
	}
}

func writeJSONErrors(w http.ResponseWriter, r *http.Request, data interface{}, errors []string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)

	if strings.ToLower(r.URL.Query().Get("envelope")) == "true" {
		e := map[string]interface{}{
			"data":   data,
			"status": code,
			"error":  errors,
		}

		writeMIME(w, r, e)
	} else {
		writeMIME(w, r, errors)
	}
}

func writeJSONError(w http.ResponseWriter, r *http.Request, data interface{}, err string, code int) {
	writeJSONErrors(w, r, data, []string{err}, code)
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
