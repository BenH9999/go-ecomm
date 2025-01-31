package main

import (
    "fmt"
    "log"
    "net/http"
)

func main(){
    fmt.Println("Starting Go server on :8080...")

    http.Handle("/", http.FileServer(http.Dir("../frontend")))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
