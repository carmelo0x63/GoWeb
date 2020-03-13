package main

import (
    "net/http"
    "os"
    "runtime"
    "strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    message := r.URL.Path
    message = strings.TrimPrefix(message, "/")
    message = "This is " + hostname + " running on " + runtime.GOOS + "/" + runtime.GOARCH + " saying: " + message + "\n"

    w.Write([]byte(message))
}

func main() {
    http.HandleFunc("/", sayHello)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

