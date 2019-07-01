package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"math/rand"
	"net/http"
	"os"
)

type Collector interface {
	Describe(chan<- *prometheus.Desc)
	Collect(chan<- prometheus.Metric)
}

type ClusterManager struct {
	Zone         string
	OOMCountDesc *prometheus.Desc
	RAMUsageDesc *prometheus.Desc
}

func (c *ClusterManager) ReallyExpensiveAssessmentOfTheSystemState() (ooMCountDesc map[string]int, ramUsageByHost map[string]float64) {
	ooMCountDesc = map[string]int{
		"192.17.81.11": int(rand.Int31n(1000)),
		"192.17.81.22": int(rand.Int31n(1000)),
	}
	ramUsageByHost = map[string]float64{
		"192.17.81.11": rand.Float64() * 100,
		"192.17.81.22": rand.Float64() * 100,
	}
	return
}

func (c *ClusterManager) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.OOMCountDesc
	ch <- c.RAMUsageDesc
}

func (c *ClusterManager) Collect(ch chan<- prometheus.Metric) {
	oomCountByHost, ramUsageByHost := c.ReallyExpensiveAssessmentOfTheSystemState()
	for host, oomCount := range oomCountByHost {
		ch <- prometheus.MustNewConstMetric(c.OOMCountDesc, prometheus.CounterValue, float64(oomCount), host)
	}
	for host, ramUsage := range ramUsageByHost {
		ch <- prometheus.MustNewConstMetric(c.RAMUsageDesc, prometheus.GaugeValue, ramUsage, host)
	}
}

func NewClusterManager(zone string) *ClusterManager {
	return &ClusterManager{
		Zone: zone,
		OOMCountDesc: prometheus.NewDesc(
			"clustermanager_oom_crashes_total",
			"Number of OOM crashes.",
			[]string{"host"},
			prometheus.Labels{"zone": zone},
		),
		RAMUsageDesc: prometheus.NewDesc(
			"clustermanager_ram_usage_bytes",
			"RAM usage as reported to the cluster manager.",
			[]string{"host"},
			prometheus.Labels{"zone": zone},
		),
	}
}

func main() {
	// cpuTemp := prometheus.NewGauge(prometheus.GaugeOpts{
	// 	Name: "cpu_temperature_celsius",
	// 	Help: "Current temperature of the CPU.",
	// })
	// hdFailures := prometheus.NewCounterVec(prometheus.CounterOpts{
	// 	Name: "hd_errors_total",
	// 	Help: "Number of hard-disk errors.",
	// }, []string{"device"})

	// prometheus.MustRegister(cpuTemp)
	// prometheus.MustRegister(hdFailures)

	// cpuTemp.Set(65.3)
	// hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
	// http.Handle("/metrics", promhttp.Handler())
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// 自定义
	workDB := NewClusterManager("db")
	workDC := NewClusterManager("dc")
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(workDB)
	reg.MustRegister(workDC)

	gatherers := prometheus.Gatherers{
		prometheus.DefaultGatherer,
		reg,
	}

	h := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{
		ErrorLog:      log.NewErrorLogger(),
		ErrorHandling: promhttp.ContinueOnError,
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Errorf("Error occur when start server %v", err)
		os.Exit(1)
	}
}
