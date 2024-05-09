package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Book struct {
    Title string `json:"title"`
}

var books []Book

func init() {
    // Populate books slice with some sample book names
    books = []Book{
        {"Book 1"},
        {"Book 2"},
        {"Book 3"},
    }

    // Write the list of book names to a file
    writeToFile("/mnt/data/books.json", books)
}

func main() {
    http.HandleFunc("/books", getBooks)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func writeToFile(filename string, data interface{}) {
	// Marshal data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling data: %v", err)
	}

	// Create or open the file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	// Write JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}
