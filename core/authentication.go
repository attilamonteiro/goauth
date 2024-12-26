package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Configurações para autenticação
const (
	TokenDuration = 72 * time.Hour // Duração do token: 72 horas
)

var (
	secretKey = []byte("SUA_CHAVE_SECRETA")
)

// GenerateToken cria um token JWT com HMAC ou RSA
func GenerateToken(userUUID string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(TokenDuration).Unix(),
		"iat": time.Now().Unix(),
		"sub": userUUID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token com HMAC
	return token.SignedString(secretKey)
}

// ValidateToken valida um token JWT e retorna as claims
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Valida o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})
}

// ExtractClaims extrai as claims de um token JWT válido
func ExtractClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token or claims structure")
	}
	return claims, nil
}

// Middleware para validar o token JWT em requisições
func RequireTokenAuthentication(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header required", http.StatusUnauthorized)
		return
	}

	// Remove prefixo "Bearer "
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	_, err := ValidateToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Continua para o próximo handler
	next(w, r)
}

