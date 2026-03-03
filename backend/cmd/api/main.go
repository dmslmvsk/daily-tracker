package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting on :8080")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})
	http.ListenAndServe(":8080", nil)
}