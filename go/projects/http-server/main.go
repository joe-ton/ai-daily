package main

import (
    "log"
    "errors"
    "net/http"
)

func generalHandle(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        w.Write([]byte("POST"))
    case http.MethodGet:
        w.Write([]byte("READ"))
    default:
        http.Error(w, "Not method", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/api/resource", generalHandle)

}
