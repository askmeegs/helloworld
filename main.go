package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello World v1"
	opsProcessed.Inc()
	w.Write([]byte(message))
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "megan_helloworld_processed_ops_total",
		Help: "The total number of processed hello calls",
	})
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("Signal received: %v", sig)
			os.Exit(0)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", sayHello)
	fmt.Println("some log message")
	fmt.Println("Hello world server listening on 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
