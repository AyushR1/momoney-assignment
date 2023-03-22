package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	// Start the server
	addr := ":8080"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running at port %s", addr)
	})
	log.Printf("Server running on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
