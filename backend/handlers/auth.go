package handlers

import (
    "encoding/json"
    "net/http"
    "go-ecomm/backend/models"
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

    row := models.DB.QueryRow("SELECT password FROM Customer WHERE email = ?", req.Email)
    var storedPassword string
    err = row.Scan(&storedPassword)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    if storedPassword != req.Password {
        http.Error (w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    response := map[string]string{"message": "Login successful"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func TestHandler(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello backend")) 
}
