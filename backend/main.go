package main

import (
    "fmt"
    "log"
    "net/http"
    "go-ecomm/backend/handlers"
    "go-ecomm/backend/models"
    //"go-ecomm/backend/middleware"
)

func main(){
    models.InitDB()
    defer models.DB.Close()

    http.HandleFunc("/api/login", handlers.LoginHandler)
    http.HandleFunc("/api/register", handlers.RegisterHandler)

    http.Handle("/", http.FileServer(http.Dir("frontend")))

    //protected := middleware.AuthMiddleware(http.FileServer(http.Dir("frontend")))
    //http.Handle("/protected/", http.StripPrefix("/protected", protected))

    fmt.Println("Starting Go server on :8000...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
