package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type HandlePostRequest struct {
    FirstName string
    LastName string
    Id int
}

// Method
func HandlePost(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Wrong Method", http.StatusMethodNotAllowed)
        return
    }

    var handlePostRequest HandlePostRequest
    err := json.NewDecoder(r.Body).Decode(&handlePostRequest)
    if err != nil {
        http.Error(w, "Wrong Payload Structure", http.StatusBadRequest)
        return
    }
    defer r.Body.Close() // clean up resources

    w.WriteHeader(http.StatusCreated) // completed
    w.Header().Set("Content-Type", "application/json") // type
    
    json.NewEncoder(w).Encode(map[string]string {
        "FirstName": "Joe",
        "LastName": "Ton",
        "Id": "1",
    })
}

type HandleReadRequest struct {
    FirstName string
    LastName string
    Id int
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Not GET Method", http.StatusMethodNotAllowed)
        return
    }

    var handleReadRequest HandleReadRequest
    err := json.NewDecoder(r.Body).Decode(&handleReadRequest)
    if err != nil {
        http.Error(w, "Not GET Payload", http.StatusBadRequest)
        return
    }

    defer r.Body.Close() // cleanup resource

    w.WriteHeader(http.StatusCreated) // crates status
    w.Header().Set("Content-Type", "application/json")

    if json.NewEncoder(w).Encode(map[string]string {
        "FirstName": "Joe",
        "LastName": "Ton",
        "Id": "1",
    })
}

func main() {
    http.HandleFunc("/api/resources", HandlePost)
    http.HandleFunc("/api/resources", HandleGet)
}


