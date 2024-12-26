package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Configurações para autenticação
const (
	TokenDuration = 72 * time.Hour // Duração do token: 72 horas
)

var (
	secretKey = []byte("SUA_CHAVE_SECRETA") // Use uma chave segura
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

// Inicializa o backend para autenticação com suporte a HMAC e RSA
func InitAuthKeys(privateKeyPath, publicKeyPath string) error {
	// Carrega a chave privada
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return fmt.Errorf("error reading private key file: %v", err)
	}
	block, _ := pem.Decode(privateKeyData)
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("error parsing private key: %v", err)
	}

	// Carrega a chave pública
	publicKeyData, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return fmt.Errorf("error reading public key file: %v", err)
	}
	block, _ = pem.Decode(publicKeyData)
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("error parsing public key: %v", err)
	}
	publicKey = pubKey.(*rsa.PublicKey)

	return nil
}

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

// HashPassword gera o hash de uma senha
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// CheckPassword compara uma senha com seu hash
func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
