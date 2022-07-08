package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	endpoint = "/metrics"
)

//Run runs metrics http server
func Run(port string, certFile string, certKey string) error {
	portNumber := ":" + port
	logrus.Infof("Starting http server to serve metrics at port '%s', endpoint '%s'", portNumber, endpoint)

	http.Handle(endpoint, promhttp.Handler())

	if certFile != "" && certKey != "" {
		// start an https server using the mux server
		return http.ListenAndServeTLS(portNumber, certFile, certKey, nil)
	} else {
		// start an http server using the mux server
		return http.ListenAndServe(portNumber, nil)
	}
}
