package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	podName := os.Getenv("HOSTNAME")
	nodeName := os.Getenv("NODE_NAME")

	response := fmt.Sprintf("Pod Name: %s\nNode Name: %s\n", podName, nodeName)
	fmt.Fprintf(w, response)
}

func main() {
	http.HandleFunc("/", handler)
	port := 8080
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server listening on port %d...\n", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
