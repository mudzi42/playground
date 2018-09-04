package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/mudzi42/playground/hydra/hlogger"
	"github.com/mudzi42/playground/hydra/shield"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web server")

	builder := shieldBuilder.NewShieldBuilder()
	shield := builder.RaiseFront().RaiseBack().Build()
	logger.Printf("%+v \n", *shield)

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8081", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprintf(w, "Welcome to the Hydra software system")
	logger.Println("Received an http Get request on root url")

}
