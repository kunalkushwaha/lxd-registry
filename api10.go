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

func getImageURL(image string) string {

	ImageMap := map[string]string{
		"ubuntu32": "http://images.linuxcontainers.org/images/ubuntu/trusty/i386/default/20141128_03:49/rootfs.tar.xz",
		"ubuntu64": "http://images.linuxcontainers.org/images/ubuntu/trusty/amd64/default/20141128_03:49/rootfs.tar.xz",
		"fedora32": "http://images.linuxcontainers.org/images/fedora/20/amd64/default/20141128_01:27/rootfs.tar.xz",
	}

	return ImageMap[image]
}

func GetImageResponse(w http.ResponseWriter, req *http.Request) {
	/*
	   TODO: Find the request in DB and redirect the request if found!!
	          else return not found.
	*/
	imagePath := getImageURL("ubuntu32")

	BuildRedirectRequest(301, imagePath, w)
}

var getImage = Command{"getImage", true, GetImageResponse, nil, nil, nil}
