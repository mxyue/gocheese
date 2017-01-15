package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/facebookgo/grace/gracehttp"
	"gocheese/apis"
	_ "gocheese/config"
	"net/http"
	"os"
)

var (
	address0 = flag.String("a0", ":48567", "Zero address to bind to.")
)

func main() {
	log.Info(fmt.Sprintf("go cheese start pid: %d", os.Getpid()))
	flag.Parse()
	gracehttp.Serve(
		&http.Server{Addr: *address0, Handler: apis.Handlers()},
	)

}
