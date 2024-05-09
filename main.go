package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Fatalf("Error marshalling data: %v", err)
    }

    err = ioutil.WriteFile(filename, jsonData, 0644)
    if err != nil {
        log.Fatalf("Error writing to file: %v", err)
    }
}
