package utils

import (
    "time"
    "crypto/rand"
    "github.com/golang-jwt/jwt"
    "golang.org/x/crypto/bcrypt"
    "github.com/google/uuid"

    "xuanqiong/config"
)

// generateRandomPassword 生成包含特殊字符的随机密码
func GenerateRandomPassword(length int64) (string, error) {
    const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/"
    password := make([]byte, length)
    _, err := rand.Read(password)
    if err != nil {
        return "", err
    }
    for i, b := range password {
        password[i] = chars[int(b)%len(chars)]
    }
    return string(password), nil
}

func GenJWTToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * time.Duration(config.Config.JWT.ExpiresIn)).Unix(),
    })
    return token.SignedString([]byte(config.Config.JWT.Secret))
}

func GenPasswordHash(password string) string {
    passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return ""
    }
    return string(passwordHash)
}

func IsHashEqual(hash string, passwd string) bool {
    if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd)); err != nil {
        return false
    }
    return true
}

func DecryptJWTToken(tokenString string) (*jwt.StandardClaims, error) {
    var claims jwt.StandardClaims
    token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.Config.JWT.Secret), nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, err
}

func GenerateUniqueID() string {
    return uuid.New().String()
}
