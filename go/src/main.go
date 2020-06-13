package main

import (
	"log"
	"net/http"

	stats_api "github.com/fukata/golang-stats-api-handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"net/http/pprof"
	_ "net/http/pprof"

	"github.com/omuricer/go-profiling/subject"
)

func handleRequests() {

	subject := subject.NewSubject()

	router := mux.NewRouter().StrictSlash(true)
	// routes about can
	router.HandleFunc("/subject/1", subject.UploadAndZip).Methods(http.MethodGet)

	// routes about stats
	router.HandleFunc("/stats", stats_api.Handler).Methods(http.MethodGet)

	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	// runtime.MemProfileRate = 1

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	handleRequests()
}
