package handlers

import (
    "encoding/json"
    "log"
    "strings"
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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"message": "Method not allowed"})
		return
	}

	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Bad request"})
		return
	}

	stmt, err := models.DB.Prepare("INSERT INTO Customer(username, email, password) VALUES(?, ?, ?)")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Username, req.Email, req.Password)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(err.Error(), "UNIQUE constraint failed: Customer.email") {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"message": "Email already registered"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "Error inserting into database"})
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer registered", "username": req.Username})
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
