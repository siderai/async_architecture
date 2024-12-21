package main

import (
	"fmt"
	"io"
	"net/http"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    
)

func asyncRequestHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}
    

	done := make(chan string)
	go makeRequest(url, done)

	result := <-done
	fmt.Fprintf(w, result)
}

func makeRequest(url string, done chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		done <- fmt.Sprintf("Error: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		done <- fmt.Sprintf("Error reading response body: %v", err)
		return
	}

	done <- string(body)
}

func main() {
	http.HandleFunc("/fetch", asyncRequestHandler)
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}