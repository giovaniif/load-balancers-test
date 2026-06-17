package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}

		fmt.Fprintf(w, "Answered by Instance: %s\n", hostname)
	})

	fmt.Println("App running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
