package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Done")
	})
	http.ListenAndServe(":8080", nil)
}
