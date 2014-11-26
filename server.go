package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	//"log"
	"fmt"
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

var api10 = []Command{
	pingStatus,
	listCmd,
}

type resp struct {
	Type     string      `json:"type"`
	Result   string      `json:"result"`
	Metadata interface{} `json:"metadata"`
}

func main() {
	d := &Daemon{}
	d.mux = mux.NewRouter()

	d.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		SyncResponse(true, "Welcome to lxd-registery!!", w)
	})

	for _, c := range api10 {
		d.createCmd("1.0", c)
	}
	
	fmt.Printf("lxd-Registery started....\n")
	fmt.Printf("Listening at <IP>:8080/\n")
	http.Handle("/", d.mux)
	// wait for clients
	http.ListenAndServe(":8080", nil)
}

func PingResponse(w http.ResponseWriter, req *http.Request) {
	SyncResponse(true, "Server Alive!", w)
}

var pingStatus = Command{"pingStatus", true, PingResponse, nil, nil, nil}

func ListCmdResponse(w http.ResponseWriter, req *http.Request) {
	SyncResponse(true, "List Cmd!", w)
}

var listCmd = Command{"listCmd", true, ListCmdResponse, nil, nil, nil}

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

func SyncResponse(success bool, metadata interface{}, w http.ResponseWriter) {
	result := "success"
	if !success {
		result = "failure"
	}

	r := resp{Type: "Resp", Result: result, Metadata: metadata}
	enc, err := json.Marshal(&r)
	if err != nil {
		//InternalError(w, err)
		return
	}
	fmt.Printf(string(enc))

	w.Write(enc)
}

func NotImplemented(w http.ResponseWriter) {
	SyncResponse(true, "Not Implemented!!", w)
}
