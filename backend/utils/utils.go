package utils

import (
    "crypto/rand"
    "encoding/base64"
    "github.com/golang-jwt/jwt"
    "golang.org/x/crypto/bcrypt"
    "github.com/google/uuid"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/mem"
)

// 生成随机Jwt secret
func GenerateRandomJwtSecret() string {
    randomBytes := make([]byte, 32)
    _, _ = rand.Read(randomBytes)
    return base64.URLEncoding.EncodeToString(randomBytes)
}

// generateRandomChars 生成包含特殊字符的随机密码
func GenerateRandomChars(length int64, chartype int64) (string, error) {
    const numChars = "0123456789"
    const specChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
    const lowerChars = "abcdefghijklmnopqrstuvwxyz"
    const upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    var chars string
    //const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/"
    randomBytes := make([]byte, length)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return "", err
    }
    // 纯数字
    if chartype == 1 {
        chars = numChars
    }
    // 纯小写字母
    if chartype == 2 {
        chars = lowerChars
    }
    // 纯大写字母
    if chartype == 3 {
        chars = upperChars
    }
    // 字母+数字
    if chartype == 4 {
        chars = numChars + lowerChars + upperChars
    }
    // 字母+数字+特殊符号
    if chartype == 5 {
        chars = numChars + lowerChars + upperChars + specChars
    }
    for i, b := range randomBytes {
        randomBytes[i] = chars[int(b)%len(chars)]
    }
    return string(randomBytes), nil
}

func GenJWTToken(username string, role int64, expires int64, secret string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "role":     role,
        "exp":      expires,
    })
    return token.SignedString([]byte(secret))
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

func DecryptJWTToken(tokenString string, secret string) (*jwt.StandardClaims, error) {
    var claims jwt.StandardClaims
    token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
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

func GetSystemUsage() (int64, int64, int64) {
    cpuPercent, _ := cpu.Percent(0, false)
    cpuUsage := int64(cpuPercent[0])

    memInfo, _ := mem.VirtualMemory()
    memUsage := int64(memInfo.UsedPercent)

    diskUsageInfo, _ := disk.Usage("/")
    diskUsage := int64(diskUsageInfo.UsedPercent)
    return cpuUsage, memUsage, diskUsage
}