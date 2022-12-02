package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(50 * time.Millisecond)

	fmt.Fprintf(w, "It's been 50ms")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting web server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
