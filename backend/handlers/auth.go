package handlers

import (
    "encoding/json"
    "net/http"
    "go-ecomm/backend/models"
    "github.com/google/uuid"
)

type RegisterRequest struct {
    Username string `json:"username"`
    Email string `json:"email"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req RegisterRequest;
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    stmt, err := models.DB.Prepare("INSERT INTO Customer(username, email, password) VALUES(?, ?, ?)")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(req.Username, req.Email, req.Password)
    if err != nil {
        http.Error(w, "Error inserting into database", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Customer registered"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req LoginRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    row := models.DB.QueryRow("SELECT username, password FROM Customer WHERE email = ?", req.Email)
    var storedUsername string
    var storedPassword string
    err = row.Scan(&storedUsername, &storedPassword)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    if storedPassword != req.Password {
        http.Error (w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    token := uuid.New().String()
    models.SetSession(token, storedUsername)
    cookie := http.Cookie{
        Name: "session_token",
        Value: token,
        Path: "/",
        HttpOnly: true,
    }
    http.SetCookie(w, &cookie)

    response := map[string]string{"message": "Login successful", "username": storedUsername}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
