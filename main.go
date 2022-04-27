package main

/*
	Atlas API Service
	Thijs Haker
*/

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Atlas-Compute-Platform/lib"
)

var (
	nsTable lib.Dict
	apiMux  *http.ServeMux
	mgrMux  *http.ServeMux
)

func usage() {
	fmt.Fprintf(os.Stderr, "Atlas API Service %s\n", lib.VERS)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var (
		apiAddr = flag.String("ap", ":80", "Specify public port")
		mgrAddr = flag.String("mp", lib.PORT, "Specify management port")
		cfgFile = flag.String("f", "config.json", "Specify config file")
		cfgBuf  []byte
		err     error
	)
	flag.Usage = usage
	flag.Parse()

	if cfgBuf, err = os.ReadFile(*cfgFile); err != nil {
		lib.LogError(os.Stderr, "main.main", err)
		os.Exit(1)
	}
	if nsTable, err = lib.ImportDict(cfgBuf); err != nil {
		lib.LogError(os.Stderr, "main.main", err)
		os.Exit(1)
	}

	mgrMux = http.NewServeMux()
	apiMux = http.NewServeMux()

	mgrMux.HandleFunc("/ping", lib.ApiPing)
	//mgrMux.HandleFunc("/list", apiList)
	//mgrMux.HandleFunc("/store", apiStore)
	//mgrMux.HandleFunc("/remove", apiRemove)
	//apiMux.HandleFunc("/", apiHandle)

	if err = http.ListenAndServe(*apiAddr, apiMux); err != nil {
		lib.LogError(os.Stderr, "main.main", err)
	}
	if err = http.ListenAndServe(*mgrAddr, mgrMux); err != nil {
		lib.LogError(os.Stderr, "main.main", err)
	}
}
