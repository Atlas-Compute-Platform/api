package main

/*
	Atlas API Service
	Thijs Haker
*/

import (
	"flag"
	"net/http"
	"os"

	"github.com/Atlas-Compute-Platform/lib"
)

var (
	nsTable lib.Dict
	apiMux  *http.ServeMux
	mgrMux  *http.ServeMux
)

func main() {
	lib.SvcName = "Atlas API Service"
	lib.SvcVers = "1.0"

	var (
		apiAddr = flag.String("ap", lib.API_PORT, "Specify public port")
		mgrAddr = flag.String("mp", lib.PORT, "Specify management port")
		cfgFile = flag.String("f", lib.CONFIG, "Specify config file")
		cfgBuf  []byte
		err     error
	)
	flag.Usage = lib.Usage
	flag.Parse()

	if cfgBuf, err = os.ReadFile(*cfgFile); err != nil {
		lib.LogFatal(os.Stderr, "main.main", err)
	}
	if nsTable, err = lib.ImportDict(cfgBuf); err != nil {
		lib.LogFatal(os.Stderr, "main.main", err)
	}

	mgrMux = http.NewServeMux()
	apiMux = http.NewServeMux()

	mgrMux.HandleFunc("/ping", lib.ApiPing)
	mgrMux.HandleFunc("/list", apiList)
	mgrMux.HandleFunc("/bind", apiBind)
	mgrMux.HandleFunc("/unbind", apiUnbind)
	apiMux.HandleFunc("/", apiHandle)

	go http.ListenAndServe(*mgrAddr, mgrMux)
	if err = http.ListenAndServe(*apiAddr, apiMux); err != nil {
		lib.LogError(os.Stderr, "main.main", err)
	}
}
