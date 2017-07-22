package uhandlers

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"ucache"
)

const (
	urlInfoPath = "/urlinfo/1/"
	lenURLInfo  = len(urlInfoPath)
	// UPDATED : status of post response
	UPDATED     = "updated"
	// ERROR : status of GET/POST response if any error
	ERROR       = "error"
	// SAFE : status of post GET response indecate url is safe
	SAFE        = "safe"
	// UNSAFE : status of post GET response indecate url is unsafe
	UNSAFE      = "unsafe"
)

func urlCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := CheckURLResponse{
		Response{Status: SAFE},
	}
	checkURL := r.URL.Path[lenURLInfo:]
	w.Header().Set("Content-Type", "application/json")

	if ucache.Get(checkURL) {
		response.Status = UNSAFE
	}
	fmt.Fprintf(w, response.JSON())
}

func urlPostHandler(w http.ResponseWriter, r *http.Request) {
	scanner := bufio.NewScanner(r.Body)
	response := UploadResponse{}
	var err error

	for scanner.Scan() {
		err = ucache.Set(scanner.Text())
		response.Count++
	}
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
	} else {
		response.Status = UPDATED
	}
	fmt.Fprintf(w, response.JSON())
}

// GlobalHandler will be mounted against "/" and all
// URL request will hit global handler and then call sub handlers
// Not used another mux libraries since we have full url as url parameter
func GlobalHandler(w http.ResponseWriter, r *http.Request) {
	urlLength := len(r.URL.Path)
	var basePath string
	if urlLength > lenURLInfo {
		basePath = r.URL.Path[:lenURLInfo]
		basePath = strings.ToLower(basePath)
	}

	if r.Method == "GET" && basePath == urlInfoPath {
		urlCheckHandler(w, r)
		return
	}
	if r.Method == "POST" && r.URL.Path == "/urls" {
		urlPostHandler(w, r)
		return
	}
}
