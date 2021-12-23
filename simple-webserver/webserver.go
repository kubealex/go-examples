package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from this awesome webserver!!")
}

func main() {
	http.HandleFunc("/", handleRoot)

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Webserver serving on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
