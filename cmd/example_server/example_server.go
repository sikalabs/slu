package example_server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ondrejsika/go-dela"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/version"
	"github.com/spf13/cobra"
)

var FlagPort int

var Cmd = &cobra.Command{
	Use:   "example-server",
	Short: "Run example web server",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		var (
			requestsTotal = prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: "example",
					Name:      "requests_total",
					Help:      "Total number of requests per endpoint",
				}, []string{"path"},
			)
		)

		prometheus.MustRegister(requestsTotal)

		portStr := strconv.Itoa(FlagPort)
		hostname, _ := os.Hostname()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			requestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server! %s %s \n", hostname, portStr)
			fmt.Printf("RemoteAddr=%s\n", r.RemoteAddr)
		})
		http.HandleFunc("/slow1s", func(w http.ResponseWriter, r *http.Request) {
			requestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			time.Sleep(1 * time.Second)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 1s)! %s %s \n", hostname, portStr)
		})
		http.HandleFunc("/slow10s", func(w http.ResponseWriter, r *http.Request) {
			requestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			time.Sleep(10 * time.Second)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 10s)! %s %s \n", hostname, portStr)
		})
		http.HandleFunc("/slow30s", func(w http.ResponseWriter, r *http.Request) {
			requestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			time.Sleep(30 * time.Second)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 30s)! %s %s \n", hostname, portStr)
		})
		http.HandleFunc("/slow60s", func(w http.ResponseWriter, r *http.Request) {
			requestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			time.Sleep(60 * time.Second)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 60s)! %s %s \n", hostname, portStr)
		})
		http.HandleFunc("/slow5m", func(w http.ResponseWriter, r *http.Request) {
			requestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			time.Sleep(5 * time.Minute)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 5m)! %s %s \n", hostname, portStr)
		})
		http.HandleFunc("/slow10m", func(w http.ResponseWriter, r *http.Request) {
			requestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			time.Sleep(10 * time.Minute)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 10m)! %s %s \n", hostname, portStr)
		})
		http.HandleFunc("/slow15m", func(w http.ResponseWriter, r *http.Request) {
			requestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			time.Sleep(15 * time.Minute)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 15m)! %s %s \n", hostname, portStr)
		})

		http.HandleFunc("/dela", func(w http.ResponseWriter, r *http.Request) {
			requestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			fmt.Fprint(w, `<!doctype html><html lang="en"><head>
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>Dela</title><style>
body, html { margin: 0; padding: 0; height: 100%; width: 100%; }
.dog-image-container { position: relative; height: 100%; width: 100%; overflow: hidden; }
.dog-image { object-fit: cover; width: 100%; height: 100%; display: block; }
</style></head><body>
<div class="dog-image-container"><img src="/dela.jpg" alt="Dela" class="dog-image" /></div>
</body></html>
			`)
		})

		http.HandleFunc("/dela.jpg", func(w http.ResponseWriter, r *http.Request) {
			requestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			w.Header().Set("Content-Type", "image/jpeg")
			w.WriteHeader(http.StatusOK)
			w.Write(dela.DELA1_JPG)
		})

		http.Handle("/metrics", promhttp.Handler())

		fmt.Println("[slu " + version.Version + "] Server started on 0.0.0.0:" + portStr + ", see http://127.0.0.1:" + portStr)
		http.ListenAndServe(":"+portStr, nil)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.PersistentFlags().IntVarP(
		&FlagPort,
		"port",
		"p",
		8000,
		"Listen on port",
	)
}
