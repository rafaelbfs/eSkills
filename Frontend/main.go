package main

import (
	"flag"
	"fmt"
	"github.com/rafaelbfs/eSkills/Frontend/rest"
	"github.com/rafaelbfs/eSkills/Frontend/security"
	"github.com/rafaelbfs/eSkills/Frontend/templates"
	"log"
	"net/http"
)

var environment string
var configFile = ""

func addPageHandlers(mux *http.ServeMux) {
	mux.HandleFunc(templates.PERSON_DETAILS_PATH, templates.HandlePersonDetails)
}

func loadCfg() {
	flag.StringVar(&environment, "env", "local",
		"The environment eg. local|test|production")
	flag.StringVar(&configFile, "cfgpath", "",
		"The location of the .env file")
	flag.Parse()

	if len(configFile) == 0 {
		configFile = fmt.Sprintf("etc/%s.env", environment)
		log.Printf("cfg file is:>%s", configFile)
	}

	security.Initialize("etc/local.env")
}

func main() {
	loadCfg()
	r := http.NewServeMux()
	fs := http.FileServer(http.Dir("./Frontend/static/")) // Frontend/rust/pkg/rust_bg.wasm
	wasm := http.FileServer(http.Dir("./Frontend/rust/pkg/"))
	addPageHandlers(r)

	r.Handle("/site/", http.StripPrefix("/site", fs))
	r.Handle("/rust/", http.StripPrefix("/rust", wasm))
	r.HandleFunc("/rest/people", rest.HandlePersonRequests)
	security.ConfigureGoogleAuth(r)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
