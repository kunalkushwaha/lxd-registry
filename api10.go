package main

import (
	"net/http"
)

var Api10 = []Command{
	pingStatus,
	listCmd,
	getImage,
}

func PingResponse(w http.ResponseWriter, req *http.Request) {
	SyncResponse(true, "Server Alive!", w)
}

var pingStatus = Command{"pingStatus", true, PingResponse, nil, nil, nil}

func ListCmdResponse(w http.ResponseWriter, req *http.Request) {
	SyncResponse(true, "List Cmd!", w)
}

var listCmd = Command{"listCmd", true, ListCmdResponse, nil, nil, nil}

func GetImageResponse(w http.ResponseWriter, req *http.Request) {
	/*
	   TODO: Find the request in DB and redirect the request if found!!
	          else return not found.
	*/
	BuildRedirectRequest(301, "http://google.com", w)
}

var getImage = Command{"getImage", true, GetImageResponse, nil, nil, nil}
