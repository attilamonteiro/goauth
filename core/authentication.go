package core

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"goauth/models"
	"goauth/settings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type JWTAuthenticationBackend struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

const (
	tokenDuration = 72 // Token dura 72 horas
	secretKey     = "your-secret-key" // Define your secret key here
)

var authBackendInstance *JWTAuthenticationBackend

// Inicializa a autenticação JWT
func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			PrivateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}
	return authBackendInstance
}

func getPrivateKey() *rsa.PrivateKey {
	privateKeyData, err := ioutil.ReadFile(settings.PrivateKeyPath)
	if err != nil {
		fmt.Println("Error reading private key file")
		os.Exit(1)
	}
	block, _ := pem.Decode(privateKeyData)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key")
		os.Exit(1)
	}
	return privateKey
}

func getPublicKey() *rsa.PublicKey {
	publicKeyData, err := ioutil.ReadFile(settings.PublicKeyPath)
	if err != nil {
		fmt.Println("Error reading public key file")
		os.Exit(1)
	}
	block, _ := pem.Decode(publicKeyData)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing public key")
		os.Exit(1)
	}
	return publicKey.(*rsa.PublicKey)
}

// Gera o token JWT
func (backend *JWTAuthenticationBackend) GenerateToken(userUUID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * tokenDuration).Unix(),
		"iat": time.Now().Unix(),
		"sub": userUUID,
	}
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Autentica o usuário
func (backend *JWTAuthenticationBackend) Authenticate(user *models.User) bool {
    // Comparação de senha com bcrypt
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
    testUser := models.User{
        Username: "haku",
        Password: string(hashedPassword),
    }

    return user.Username == testUser.Username &&
        bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
}

