package urlshort

import "net/http"

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	Queries     []string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

var routes = Routes{
	Route{
		"Hello",
		"GET",
		"/",
		nil,
		hello,
	},
}
