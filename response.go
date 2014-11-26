package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type resp struct {
	Type     string      `json:"type"`
	Result   string      `json:"result"`
	Metadata interface{} `json:"metadata"`
}

type Jmap map[string]interface{}

func NotImplemented(w http.ResponseWriter) {
	ErrorResponse(501, "not implemented", w)
}

func BuildRedirectRequest(code int, url string, w http.ResponseWriter) {
	w.Header().Set("Location", url)
	w.WriteHeader(code)
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
	fmt.Printf(string(enc) + "\n")

	w.Write(enc)
}

func ErrorResponse(code int, msg string, w http.ResponseWriter) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(Jmap{"type": "Error", "error": msg, "error_code": code})

	if err != nil {
		// Can't use InternalError here
		http.Error(w, "Error encoding error response!", 500)
		return
	}

	http.Error(w, buf.String(), code)
}
