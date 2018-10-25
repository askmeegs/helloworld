package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello World v1"
	w.Write([]byte(message))
}
func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("Signal received: %v", sig)
			os.Exit(0)
		}
	}()

	http.HandleFunc("/", sayHello)
	fmt.Println("some log message")
	fmt.Println("Hello world server listening on 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
