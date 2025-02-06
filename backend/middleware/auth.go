package middleware

import(
    "net/http"
    "go-ecomm/backend/models"
)

func AuthMiddleware(next http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){
        cookie, err := r.Cookie("session_token")
        if err != nil{
            http.Error(w, "Unauthorised", http.StatusUnauthorized)
            return
        }

        _, ok := models.GetSession(cookie.Value)
        if !ok {
            http.Error(w, "Unauthorised", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}
