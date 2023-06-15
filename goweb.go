package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
)

var listenPort = ":8080"

func sayHello(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "This is " + hostname +
                  " running on " + runtime.GOOS +
                  "/" + runtime.GOARCH + " saying: " +
                  message + "\n"
	fmt.Printf("GoWeb: got / request\n")
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(listenPort, nil); err != nil {
		panic(err)
	}
}
