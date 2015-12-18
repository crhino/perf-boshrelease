package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/crhino/perf-server/handlers"
	"github.com/gorilla/mux"
)

var cmdPath = flag.String("cmd", "", "Path of the shell script to run")

func main() {
	flag.Parse()
	r := mux.NewRouter()
	r.Handle("/flamegraph", handlers.NewFlameGraphHandler(*cmdPath))
	http.Handle("/", r)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
