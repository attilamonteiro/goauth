package core

import (
    "net/http"
    "strings"
    "goauth/utils"
)

func RequireTokenAuthentication(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Authorization header required", http.StatusUnauthorized)
        return
    }

    // Remove prefixo "Bearer "
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")

    token, err := utils.ValidateToken(tokenString)
    if err != nil || !token.Valid {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    next(w, r)
}
