package main

import (
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Atlas-Compute-Platform/lib"
)

func apiList(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var err error

	if err = lib.SendDict(nsTable, w); err != nil {
		lib.LogError(os.Stderr, "main.apiList", err)
	}
}

func apiHandle(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var (
		val string
		ok  bool
		err error
		px  *url.URL
		pr  *http.Response
	)
	if val, ok = nsTable[r.URL.Path]; !ok {
		http.NotFound(w, r)
		return
	}
	if px, err = url.Parse(val); err != nil {
		lib.LogError(w, "main.apiHandle", err)
		return
	}

	r.Host = px.Host
	r.URL.Host = px.Host
	r.URL.Path = px.Path
	r.URL.Scheme = px.Scheme
	r.RequestURI = ""

	if pr, err = http.DefaultClient.Do(r); err != nil {
		lib.LogError(os.Stderr, "main.apiHandle", err)
		lib.LogError(w, "main.apiHandle", err)
		return
	}
	io.Copy(w, pr.Body)
}
