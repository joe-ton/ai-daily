package main

import (
    "net/http"
    "encoding/json"
)

// JSON Payload
type RequestPayload struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

// handle POST method
func createResource(w http.ResponseWriter, r *http.Request) {
    // Not POST, 405
    if r.Method != http.MethodPost {
        http.Error(w, "Not POST Method", http.StatusMethodNotAllowed)
        return
    }

    // Bad Payload, 400
    var requestPayload RequestPayload
    // need to pass by reference, not value
    err := json.NewDecoder(r.Body).Decode(&requestPayload)
    
    if err != nil {
        http.Error(w, "error payload", http.StatusBadRequest)
        return
    }

    defer r.Body.Close() // resource allocation, cleanup

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated) // 201

    json.NewEncoder(w).Encode(map[string]string {
        "name": requestPayload.Name,
        "email": requestPayload.Email,
    })

}

func main() {
    http.HandleFunc("/resource", createResource)
    http.ListenAndServe(":8080", nil)
}
