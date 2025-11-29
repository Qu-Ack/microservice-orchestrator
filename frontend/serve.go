package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
