package main

import (
    "fmt"
    "log"
    "net/http"
    "go-ecomm/backend/handlers"
    "go-ecomm/backend/models"
)

func main(){
    models.InitDB()
    defer models.DB.Close()


    http.HandleFunc("/api/test", handlers.TestHandler)
    http.HandleFunc("/api/login", handlers.LoginHandler)
    http.Handle("/", http.FileServer(http.Dir("frontend")))

    fmt.Println("Starting Go server on :8000...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
