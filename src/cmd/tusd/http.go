package main

import (
	"fmt"
	"net/http"
	"log"
)

func serveHttp() error {
	http.HandleFunc("/", route)

	addr := ":1080"
	log.Printf("serving clients at %s", addr)

	return http.ListenAndServe(addr, nil)
}

func route(w http.ResponseWriter, r *http.Request) {
	log.Printf("request: %s %s", r.Method, r.URL.RequestURI())

	w.Header().Set("Server", "tusd")

	if r.Method == "POST" && r.URL.Path == "/files" {
		createFile(w, r)
	} else {
		reply(w, http.StatusNotFound, "No matching route")
	}
}

func reply(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	fmt.Fprintf(w, "%d - %s: %s\n", code, http.StatusText(code), message)
}

func createFile(w http.ResponseWriter, r *http.Request) {
	contentRange, err := parseContentRange(r.Header.Get("Content-Range"))
	if err != nil {
		reply(w, http.StatusBadRequest, err.Error())
		return
	}

	if contentRange.Size == -1 {
		reply(w, http.StatusBadRequest, "Content-Range must indicate total file size.")
		return
	}

	if contentRange.End != -1 {
		reply(w, http.StatusNotImplemented, "File data in initial request.")
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_ = contentType
	id := uid()
	
	w.Header().Set("Location", "/files/"+id)
}