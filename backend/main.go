package main

import (
    "fmt"
    "log"
    "net/http"
    "go-ecomm/backend/handlers"
)

func main(){
    fmt.Println("Starting Go server on :8000...")

    http.HandleFunc("/api/test", handlers.TestHandler)

    http.Handle("/", http.FileServer(http.Dir("frontend")))

    log.Fatal(http.ListenAndServe(":8000", nil))
}
