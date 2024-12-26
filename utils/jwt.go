package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("SUA_CHAVE_SECRETA")

func GenerateToken(userUUID string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
    claims["iat"] = time.Now().Unix()
    claims["sub"] = userUUID

    return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return secretKey, nil
    })
}
