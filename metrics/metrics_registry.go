// Package metrics provides measures of quantitative assessment commonly used for comparing, and tracking performance or production
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
	RazorRegistry.MustRegister(ClientMetric)
}
