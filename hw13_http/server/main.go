package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	var ip, port string
	flag.StringVar(&ip, "ip", "127.0.0.1", "IP address")
	flag.StringVar(&port, "port", "8080", "Port number")
	flag.Parse()

	server := &http.Server{
		Addr:         ip + ":" + port,
		Handler:      http.DefaultServeMux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	http.HandleFunc("/", handler)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Receive %s request for %s\n", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("This is GET request!"))
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		w.Write([]byte(fmt.Sprintf("Receive %s request with body %s\n", r.Method, body)))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
