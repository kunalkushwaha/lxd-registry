package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// A Daemon can respond to requests from a lxd client.
type Daemon struct {
	mux *mux.Router
}

type Command struct {
	name         string
	untrustedGet bool
	GET          func(w http.ResponseWriter, r *http.Request)
	PUT          func(w http.ResponseWriter, r *http.Request)
	POST         func(w http.ResponseWriter, r *http.Request)
	DELETE       func(w http.ResponseWriter, r *http.Request)
}

func main() {
	d := &Daemon{}
	d.mux = mux.NewRouter()

	d.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		SyncResponse(true, "Welcome to lxd-registery!!", w)
	})

	for _, c := range Api10 {
		d.createCmd("1.0", c)
	}

	fmt.Printf("lxd-Registery started....\n")
	fmt.Printf("Listening at <IP>:8080/\n")
	http.Handle("/", d.mux)
	// wait for clients
	http.ListenAndServe(":8080", nil)
}

func (d *Daemon) createCmd(version string, c Command) {
	var uri string
	if c.name == "" {
		uri = fmt.Sprintf("/%s", version)
	} else {
		uri = fmt.Sprintf("/%s/%s", version, c.name)
	}

	d.mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			if c.GET == nil {
				NotImplemented(w)
			} else {
				c.GET(w, r)
			}
		case "PUT":
			if c.PUT == nil {
				NotImplemented(w)
			} else {
				c.PUT(w, r)
			}
		case "POST":
			if c.POST == nil {
				NotImplemented(w)
			} else {
				c.POST(w, r)
			}
		case "DELETE":
			if c.DELETE == nil {
				NotImplemented(w)
			} else {
				c.DELETE(w, r)
			}
		}
	})
}
