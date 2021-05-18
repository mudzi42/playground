package urlshort

import (
	"fmt"
	"log"
	"net/http"

	yamlV2 "gopkg.in/yaml.v2"
)

// Exception defines the exception message
// type Exception struct {
// 	Message string `json:"message"`
// }

// notFound responds with a 404
func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Not found: '%s'", r.RequestURI), http.StatusNotFound)
}

// SendError reponds with an error message and error code.
// func SendError(w http.ResponseWriter, err error) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusBadRequest)
// 	json.NewEncoder(w).Encode(Exception{Message: err.Error()})
// }

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, world!")
}

// Build the MapHandler using the mux as the fallback
// 	pathsToUrls := map[string]string{
// 		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
// 		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
// 	}
// mapHandler := urlshort.MapHandler(pathsToUrls, mux)

// 	// Build the YAMLHandler using the mapHandler as the
// 	// fallback
// 	yaml := `
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution
// `
// 	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
// 	if err != nil {
// 		panic(err)
// 	}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("looking for %s in map", r.URL.Path)

		if path, ok := pathsToUrls[r.URL.Path]; ok {
			log.Printf("redirecting %s to %s", r.URL.Path, path)
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			log.Printf("fallback from map")
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// var yamlURLs []struct {
	// 	Path string `yaml:"path"`
	// 	URL  string `yaml:"url"`
	// }

	log.Printf("looking for %s in yaml", yaml)
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	// err := yamlV2.Unmarshal(yaml, &yamlURLs)
	// if err != nil {
	// 	return nil, err
	// }

	//pathMap := make(map[string]string)
	pathMap := buildMap(parsedYaml)
	log.Printf("fallback from yaml")

	return MapHandler(pathMap, fallback), nil

}

func parseYAML(yaml []byte) (dst []map[string]string, err error) {
	err = yamlV2.Unmarshal(yaml, &dst)
	return dst, err
}

func buildMap(parsedYaml []map[string]string) map[string]string {
	mergedMap := make(map[string]string)
	for _, entry := range parsedYaml {
		key := entry["path"]
		mergedMap[key] = entry["url"]
	}
	return mergedMap
}
