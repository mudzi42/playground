package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"git.soma.salesforce.com/chudgins/gophercises/urlshort"
)

func main() {
	address := flag.String("address", "localhost", "httpd address")
	port := flag.String("port", "8080", "httpd port")
	yamlFile := flag.String("yamlFile", "", "Yaml file of URL redirects")
	flag.Parse()

	fmt.Println(*yamlFile) // not yet implemented
	router := urlshort.NewRouter()

	// Build the MapHandler using the mux router as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, router)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `	
	- path: /urlshort
	 url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
	 url: https://github.com/gophercises/urlshort/tree/solution
	`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	log.Printf("Gophercise Ex 2 - URL Shortner\nStarting the server on %s:%s\n", *address, *port)
	http.ListenAndServe(*address+":"+*port, yamlHandler)

}
