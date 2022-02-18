package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/razor-network/goInfo"
	"razor/core"
	"runtime"
)

var (
	RazorRegistry *prometheus.Registry

	// ErrorsMetric metric shows total errors count
	// MUST contain a label with a "error" key
	ErrorsMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "errors_total",
			Help: "Total amount of errors",
		},
		[]string{"error"},
	)

	osInfo = goInfo.GetInfo()

	ClientMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "client_information",
		Help: "Hold the information of client",
		ConstLabels: map[string]string{
			"Core":             osInfo.Core,
			"Platform":         osInfo.Platform,
			"razor_go_version": core.VersionWithMeta,
			"go_version":       runtime.Version(),
		},
	})
)

func init() {
	//create a registry
	RazorRegistry = prometheus.NewRegistry()

	//register razor metrics into registry

	RazorRegistry.MustRegister(ErrorsMetric)
	RazorRegistry.MustRegister(ClientMetric)
}

// ErrorInc calls prometheus counter which counts amount of errors,
// occurred on error, fatal and panic log levels
func ErrorInc() {
	ErrorsMetric.With(prometheus.Labels{"error": "error"}).Inc()
}
