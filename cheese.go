package main

import (
	"flag"
	"github.com/facebookgo/grace/gracehttp"
	"gocheese/apis"
	"net/http"
)

var (
	address0 = flag.String("a0", ":48567", "Zero address to bind to.")
)

func main() {
	flag.Parse()
	gracehttp.Serve(
		&http.Server{Addr: *address0, Handler: apis.Handlers()},
	)
}
