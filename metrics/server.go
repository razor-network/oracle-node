package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	port     = ":2112"
	endpoint = "/metrics"
)

//Run runs metrics http server
func Run() error {
	logrus.Infof("Starting http server to serve metrics at port '%s', endpoint '%s'", port, endpoint)
	server := http.NewServeMux()
	//server.Handle(endpoint, promhttp.HandlerFor(HealthRegistry, promhttp.HandlerOpts{}))

	http.Handle(endpoint, promhttp.Handler())

	// start an http server using the mux server
	return http.ListenAndServe(port, server)
}
