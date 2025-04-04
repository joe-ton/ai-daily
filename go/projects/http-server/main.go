package main

import (
    "log"
    "net/http"
)



func resourceHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {

    }

}

func main() {
    // register main route
    http.HandleFunc("/api/resource", resourceHandler)
}
