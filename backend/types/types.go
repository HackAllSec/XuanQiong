package types

import (
    "time"

    "github.com/dgrijalva/jwt-go"
)

// User 用户表
type User struct {
    ID         uint64      `gorm:"primaryKey"`
    Username   string      `gorm:"unique"`
    Password   string
    Role       int64
    CreateTime time.Time
    Status     int64
    Token      string
}

// Vulnerability 漏洞表
type Vulnerability struct {
    ID             string    `gorm:"primaryKey"`
    CVE            string
    CNNVD          string
    CNVD           string
    VulnName       string
    VulnType       string
    VulnLevel      string
    Description    string
    AffectedProduct string
    Poc            string
    Exp            string
    RepairSuggestion string
    Submit         string
    CreateTime     time.Time
    UpdateTime     time.Time
}

// Lockip 锁定IP表
type Lockip struct {
    ID         uint64    `gorm:"primaryKey"`
    ClientIP   string
    LockoutUntil  *time.Time
}

type Jwt struct {
    Role       uint64    `json:"role"`
    Username   string    `json:"username"`
    jwt.StandardClaims
}

type LoginData struct {
    Username string    `json:"username"`
    Password string    `json:"password"`
}
