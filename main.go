package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var BASE_PORT = ":9090" // means localhost:8080

func main() {
	fmt.Println("-----------------------------")
	fmt.Println("Server Web Application")

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = BASE_PORT
	} else {
		httpPort = ":" + httpPort
	}

	fmt.Println("Starting Server on port " + httpPort)

	http.HandleFunc("/server", sendHostName)
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatalf("Failed to start server on %s: %v", httpPort, err)
	}
}

func sendHostName(w http.ResponseWriter, req *http.Request) {
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	hostName += "\n"

	w.Header().Add(http.CanonicalHeaderKey("content-type"),
		"text/plain")
	w.Write([]byte(hostName))

}
