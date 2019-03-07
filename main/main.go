package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"urlshort"
)

func main() {
	yamlFile := flag.String("file", "default.yaml", `a yaml file in the format of 
	- path: /some-path-name
	  url: https://some-url.com/demo
	`)
	flag.Parse()

	yamlBytes, err := ioutil.ReadFile(*yamlFile)
	check(err)

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://github.com/JPTredway/urlshort",
		"/yaml-godoc":     "https://github.com/JPTredway/go-quiz",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler(yamlBytes, mapHandler)
	check(err)

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	return mux
}
